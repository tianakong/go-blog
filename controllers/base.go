package controllers

import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare()  {
	userId := this.GetSession("userId")
	userName := this.GetSession("userName")

	this.Data["userId"] = userId
	this.Data["userName"] = userName
	this.Data["test"] = "100"
}

func (this *BaseController) ajaxReturn(status int, info string) {
	res := make(map[string]interface{})
	res["status"] = status
	res["info"] = info
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}


func (this *BaseController) md5(pass string) string {
	h := md5.New()
	h.Write([]byte(pass))
	str := h.Sum(nil)
	return hex.EncodeToString(str)
}