package controllers

type PostsController struct {
	BaseController
}

func (this *PostsController) Add() {
	this.TplName = "posts/add.tpl"
}
