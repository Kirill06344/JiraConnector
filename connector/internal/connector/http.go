package connector

import (
	"errors"
	"fmt"
	"github.com/stewie/config"
	"net/http"
	"time"
)

var (
	delayer         <-chan time.Time
	currentWaitTime time.Duration
)

func HTTPRequest(url string) (*http.Response, error) {
	cfg, _ := config.Load()
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
