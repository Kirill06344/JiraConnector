package entity

type Project struct {
	Id    uint
	Title string
	Key   string
}

func (Project) TableName() string {
	return "project"
}
