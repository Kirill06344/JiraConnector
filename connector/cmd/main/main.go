package main

import (
	"fmt"
	"github.com/stewie/config"
	"github.com/stewie/internal/router"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Program.Port),
		Handler: router.NewRouter(),
	}

	err = srv.ListenAndServe()
}

//curl https://issues.apache.org/jira/rest/api/2/project -- all projects
//curl https://issues.apache.org/jira/rest/api/2/project/{key} -- get project by key
//curl https://issues.apache.org/jira/rest/api/2/search?jql=project=AAR&maxResults=50&expand=changelog
