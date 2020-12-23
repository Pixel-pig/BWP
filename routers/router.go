package routers

import (
	"BWP/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//登录页面
    beego.Router("/", &controllers.MainController{})
    //登录操作
	beego.Router("/Userlogin", &controllers.UserloginController{})
}
