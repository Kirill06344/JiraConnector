package controller

import "backend/internal/service"

type Group struct {
	Project *Project
	Issue   *Issue
}

func NewGroup(service *service.Service) *Group {
	return &Group{
		Project: NewProjectController(service.Project),
	}
}
