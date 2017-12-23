package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	Username   string `orm:"size(20);unique"`
	Password   string `orm:"size(32)"`
	Createtime time.Time
	Logintime  time.Time
}

func (u *User) TableName() string {
	return TableName("user")
}


func UserAdd(u *User) (int64, error) {
	u.Createtime = time.Now()
	u.Logintime = time.Now()
	return orm.NewOrm().Insert(u)
}


func UserGetByName(username string) (*User, error) {
	user := new(User)
	err := orm.NewOrm().QueryTable(user).Filter("Username", username).One(user)
	if err == nil {
		return user,nil
	}
	return nil,err
}