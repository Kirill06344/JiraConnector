package entity

type Project struct {
	ID    uint
	Title string
	Key   string
}

func (Project) TableName() string {
	return "project"
}
