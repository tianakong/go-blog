package models

import "time"

type Posts struct {
	Id         int
	Cid        int
	Uid        int
	Username   string `orm:"size(20)"`
	Title      string `orm:"size(100)"`
	Content    string `orm:"type(text)"`
	CreateTime time.Time
	ViewNumber int
}

func (p *Posts) TableName() string {
	return TableName("posts")
}
