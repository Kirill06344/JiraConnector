package entity

type Author struct {
	ID   uint
	Name string
}

func (Author) TableName() string {
	return "author"
}
