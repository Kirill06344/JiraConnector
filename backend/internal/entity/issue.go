package entity

import "time"

type Issue struct {
	ID          uint
	ProjectId   uint
	AuthorId    uint
	AssigneeId  uint
	Key         string
	CreatedTime time.Time
	ClosedTime  time.Time
	UpdatedTime time.Time
	Summary     string
	Description string
	Priority    string
	Status      string
}
