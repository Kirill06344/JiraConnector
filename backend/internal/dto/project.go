package dto

type Project struct {
	Id          uint   `json:"Id"`
	URL         string `json:"Url,omitempty"`
	Description string `json:"Description,omitempty"`
	Key         string `json:"Key,omitempty"`
	Name        string `json:"Name,omitempty"`
}
