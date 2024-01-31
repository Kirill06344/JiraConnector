package connector

import (
	"errors"
	"fmt"
	"github.com/stewie/internal/application"
	"net/http"
	"time"
)

var (
	delayer         <-chan time.Time
	currentWaitTime time.Duration
)

func GetJiraConnection(url string) (*http.Response, error) {
	cfg := application.App.Config()
	delayer = time.Tick(cfg.Program.MinTimeSleep)
	response, err := http.Get(url)
	if err == nil {
		return response, nil
	}
	currentWaitTime = cfg.Program.MinTimeSleep * 2
	for currentWaitTime <= cfg.Program.MaxTimeSleep {
		response, err = delayedRequest(url)
		if err != nil {
			return response, nil
		}
	}
	return nil, errors.New("waiting limit exceeded")
}

func delayedRequest(url string) (*http.Response, error) {
	<-delayer
	fmt.Println("Delayed request!")
	response, err := http.Get(url)
	if err != nil {
		return response, nil
	}
	currentWaitTime *= 2
	delayer = time.Tick(currentWaitTime)
	return nil, err
}

func jiraUrlAllProjects() string {
	cfg := application.App.Config()
	return cfg.Program.JiraUrl + "project"
}

func jiraUrlProjectWithKey(key string) string {
	cfg := application.App.Config()
	return cfg.Program.JiraUrl + "project/" + key
}

func jiraUrlIssuesInfo(name string) string {
	cfg := application.App.Config()
	return fmt.Sprintf("%ssearch?jql=project=%s", cfg.Program.JiraUrl, name)
}

func jiraUrlIssues(name string, startedAt int) string {
	cfg := application.App.Config()
	return fmt.Sprintf("%ssearch?jql=project=%s&expand=changelog&startAt=%d&maxResults=%d",
		cfg.Program.JiraUrl,
		name,
		startedAt,
		int(cfg.Program.IssueInOneRequest),
	)
}
