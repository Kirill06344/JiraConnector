package connector

import (
	"encoding/json"
	"fmt"
	"github.com/stewie/internal/application"
	"github.com/stewie/internal/dto"
	"github.com/stewie/internal/pb"
	"github.com/stewie/internal/utils"
	"io"
	"strings"
	"sync"
)

var mutex sync.Mutex

func GetProjects(params *utils.PageParams) (*pb.AllProjectResponse, error) {
	response, err := GetJiraConnection(jiraUrlAllProjects())

	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var projects []dto.Project
	_ = json.Unmarshal(body, &projects)

	pageInfo := &dto.PageInfo{}
	projects = utils.Filter(projects, func(project dto.Project) bool {
		return strings.HasPrefix(strings.ToLower(project.Name), strings.ToLower(params.Search)) ||
			strings.HasPrefix(strings.ToLower(project.Key), strings.ToLower(params.Search))
	})

	pageInfo.PageCount = len(projects) / params.Limit
	if len(projects)%params.Limit != 0 {
		pageInfo.PageCount++
	}
	pageInfo.ProjectsCount = len(projects)

	if params.Page*params.Limit < len(projects) {
		projects = projects[(params.Page-1)*params.Limit : params.Page*params.Limit]
		pageInfo.CurrentPage = params.Page
	} else {
		projects = projects[(params.Page-1)*params.Limit:]
		pageInfo.CurrentPage = 1
	}

	result := &pb.AllProjectResponse{
		Projects: make([]*pb.Project, len(projects)),
		PageInfo: &pb.PageInfo{
			PageCount:     int32(pageInfo.PageCount),
			CurrentPage:   int32(pageInfo.CurrentPage),
			ProjectsCount: int32(pageInfo.ProjectsCount),
		},
	}

	for i, project := range projects {
		result.Projects[i] = &pb.Project{
			Key:         project.Key,
			Name:        project.Name,
			Url:         project.URL,
			Description: project.Description,
		}
	}

	return result, nil
}

func DownloadProject(key string) (uint, error) {
	response, err := GetJiraConnection(jiraUrlProjectWithKey(key))

	if err != nil {
		return 0, err
	}
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var project dto.Project
	_ = json.Unmarshal(body, &project)

	issues, err := downloadIssues(&project)
	if err != nil {
		return 0, err
	}

	db := application.App.DB()
	return db.InsertData(&project, issues)
}

func downloadIssues(project *dto.Project) ([]dto.Issue, error) {
	cfg := application.App.Config()
	response, err := GetJiraConnection(jiraUrlIssuesInfo(project.Name))
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var issuesInfo dto.IssuesInfo
	_ = json.Unmarshal(body, &issuesInfo)

	threadCount := int(cfg.Program.ThreadCount)
	parallelDownloadsAmount := issuesInfo.Total/(int(cfg.Program.IssueInOneRequest)*threadCount) + 1
	var issues []dto.Issue
	doneRequestCount := 0
	for i := 0; i < parallelDownloadsAmount; i++ {
		var waitGroup sync.WaitGroup
		waitGroup.Add(threadCount)
		for j := 0; j < threadCount; j++ {
			issueIndex := int(cfg.Program.IssueInOneRequest) * (i*threadCount + j)
			go func() {
				defer waitGroup.Done()
				response, err = GetJiraConnection(jiraUrlIssues(project.Name, issueIndex))
				if err != nil {
					return
				}

				body, _ := io.ReadAll(response.Body)

				var newIssues dto.Issues
				err = json.Unmarshal(body, &newIssues)
				if err != nil {
					fmt.Println(err)
					return
				}
				if len(newIssues.Data) == 0 {
					return
				}

				mutex.Lock()
				defer mutex.Unlock()
				issues = append(issues, newIssues.Data...)
				doneRequestCount += len(newIssues.Data)
				fmt.Println(len(issues))
			}()
		}
		waitGroup.Wait()
		if err != nil {
			return nil, err
		}
	}
	return issues, nil
}
