package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.Router("/register", &controllers.UserController{}, "*:Register")
	beego.Router("/login", &controllers.UserController{}, "*:Login")
	beego.Router("/loginout", &controllers.UserController{}, "*:Loginout")
	beego.Router("/posts/add", &controllers.PostsController{}, "*:Add")

}
