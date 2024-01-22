package connector

import (
	"encoding/json"
	"github.com/stewie/internal/dto"
	"github.com/stewie/internal/utils"
	"io"
	"strings"
)

func GetProjects(params *utils.PageParams) (utils.Envelope, error) {
	response, err := HTTPRequest("https://issues.apache.org/jira/rest/api/2/project")

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
