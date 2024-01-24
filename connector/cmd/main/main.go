package main

import (
	"fmt"
	"github.com/stewie/internal/connector"
)

func main() {
	err := connector.DownloadProject("AAR")
	fmt.Println(err)
	//err := application.Configure()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//srv := &http.Server{
	//	Addr:    fmt.Sprintf(":%d", application.App.Config().Program.Port),
	//	Handler: router.NewRouter(),
	//}
	//err = srv.ListenAndServe()
}
