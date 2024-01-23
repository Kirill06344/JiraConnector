package connector

import (
	"encoding/json"
	"fmt"
	"github.com/stewie/internal/dto"
	"github.com/stewie/internal/utils"
	"io"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var mutex sync.Mutex

func GetProjects(params *utils.PageParams) (utils.Envelope, error) {
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

	if params.Page*params.Limit > len(projects) {
		projects = projects[(params.Page-1)*params.Limit:]
		pageInfo.CurrentPage = params.Page
	} else {
		projects = projects[(params.Page-1)*params.Limit : params.Page*params.Limit]
		pageInfo.CurrentPage = 1
	}

	env := utils.Envelope{
		"data":     projects,
		"pageInfo": pageInfo,
	}
	return env, nil
}

func DownloadProject(name string, key string) error {
	response, err := GetJiraConnection(jiraUrlProjectWithKey(key))

	if err != nil {
		return err
	}
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var project dto.Project
	_ = json.Unmarshal(body, &project)

	issues, err := downloadIssues(name)
	if err != nil {
		return err
	}
	println(len(issues))
	//to db
	return nil
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func downloadIssues(name string) ([]dto.Issue, error) {
	response, err := GetJiraConnection(jiraUrlIssuesInfo(name))
	if err != nil {
		return nil, err
	}
	body, _ := io.ReadAll(response.Body)
	defer response.Body.Close()

	var issuesInfo dto.IssuesInfo
	_ = json.Unmarshal(body, &issuesInfo)

	threadCount := int(cfg.Program.ThreadCount)
	//fix this
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
				response, err = GetJiraConnection(jiraUrlIssues(name, issueIndex))
				if err != nil {
					return
				}

				body, _ := io.ReadAll(response.Body)

				var newIssues dto.Issues
				_ = json.Unmarshal(body, &newIssues)
				if len(newIssues.Data) == 0 {
					return
				}

				mutex.Lock()
				defer mutex.Unlock()
				issues = append(issues, newIssues.Data...)
				doneRequestCount += len(newIssues.Data)
			}()
		}
		waitGroup.Wait()
		if err != nil {
			return nil, err
		}
	}
	return issues, nil
}
