package models

type Category struct {
	Id     int
	Title  string `orm:"size(10)"`
	Status int8
	Uid    int
}

func (c *Category) TableName() string {
	return TableName("category");
}
