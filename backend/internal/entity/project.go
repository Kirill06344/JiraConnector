package entity

type Project struct {
	Id    uint
	Title string
}

func (Project) TableName() string {
	return "project"
}
