package entity

type Project struct {
	Id    int64
	Title string
}

func (Project) TableName() string {
	return "project"
}
