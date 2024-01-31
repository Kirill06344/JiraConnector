package dto

type PageInfo struct {
	PageCount     int `json:"pageCount"`
	CurrentPage   int `json:"currentPage"`
	ProjectsCount int `json:"projectsCount"`
}
