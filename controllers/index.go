package controllers

import (
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	this.TplName = "index.tpl"
}
