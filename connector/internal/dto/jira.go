package dto

type Project struct {
	ID          uint
	Key         string `json:"key"`
	Name        string `json:"name"`
	URL         string `json:"self"`
	Description string `json:"description"`
}

type IssuesInfo struct {
	StartAt    int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
}

type Issues struct {
	MaxResults int     `json:"maxResults"`
	Data       []Issue `json:"issues"`
}

type Issue struct {
	Key       string      `json:"key"`
	Fields    IssueFields `json:"fields"`
	ChangeLog ChangeLog   `json:"changelog"`
}

type IssueFields struct {
	CreatedTime string        `json:"created"`
	UpdatedTime string        `json:"updated"`
	Description string        `json:"description"`
	Summary     string        `json:"summary"`
	Creator     Author        `json:"creator"`
	Assignee    Author        `json:"reporter"`
	TimeSpent   int           `json:"timespent"`
	Type        IssueType     `json:"issuetype"`
	Status      IssueStatus   `json:"status"`
	Priority    IssuePriority `json:"priority"`
}

type Author struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type IssueType struct {
	Name string `json:"name"`
}

type IssueStatus struct {
	Name string `json:"name"`
}

type IssuePriority struct {
	Name string `json:"name"`
}

type ChangeLog struct {
	Count     int       `json:"total"`
	Histories []History `json:"histories"`
}

type History struct {
	Author      Author         `json:"author"`
	CreatedTime string         `json:"created"`
	Items       []StatusChange `json:"items"`
}

type StatusChange struct {
	FromStatus string `json:"fromString"`
	ToStatus   string `json:"toString"`
}
