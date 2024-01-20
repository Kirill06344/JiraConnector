package response

import (
	"time"
)

type Issue struct {
	Id          int64     `json:"Id"`
	Project     Project   `json:"Project"`
	Key         string    `json:"Key"`
	CreatedTime time.Time `json:"CreatedTime"`
	ClosedTime  time.Time `json:"ClosedTime"`
	UpdatedTime time.Time `json:"UpdatedTime"`
	Summary     string    `json:"Summary"`
	Description string    `json:"Description"`
	Priority    string    `json:"Priority"`
	Status      string    `json:"Status"`
}
