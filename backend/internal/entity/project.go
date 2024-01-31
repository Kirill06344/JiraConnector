package entity

type Project struct {
	ID    uint
	Title string
	Key   string
	Name  string
	Url   string
}

func (Project) TableName() string {
	return "project"
}
