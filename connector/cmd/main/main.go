package main

import (
	"fmt"
	"github.com/stewie/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}

//curl https://issues.apache.org/jira/rest/api/2/project -- all projects
//curl https://issues.apache.org/jira/rest/api/2/project/{key} -- get project by key
//curl https://issues.apache.org/jira/rest/api/2/search?jql=project=AAR&maxResults=50&expand=changelog
