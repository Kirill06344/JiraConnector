package dto

import (
	"time"
)

type Issue struct {
	Id          uint      `json:"Id"`
	Project     Project   `json:"Project"`
	Key         string    `json:"Key"`
	CreatedTime time.Time `json:"CreatedTime"`
	ClosedTime  time.Time `json:"ClosedTime"`
	UpdatedTime time.Time `json:"UpdatedTime"`
	Summary     string    `json:"Summary"`
	Description string    `json:"Description"`
	Priority    string    `json:"Priority"`
	Creator     string    `json:"Creator"`
	Assignee    string    `json:"Assignee"`
	Status      string    `json:"Status"`
}
