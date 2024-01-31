package controller

import "backend/internal/service"

type Group struct {
	Project   *Project
	Issue     *Issue
	Connector *Connector
}

func NewGroup(service *service.Service) *Group {
	return &Group{
		Project:   NewProjectController(service.Project),
		Issue:     NewIssueController(service.Issue),
		Connector: NewConnectorController(service.Connector),
	}
}
