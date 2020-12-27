package routers

import (
	"BWP/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//主页面
    beego.Router("/", &controllers.MainController{})
    //注册页面
    beego.Router("/register.html", &controllers.RegisterController{})
    //登入页面
    beego.Router("/login.html", &controllers.LoginController{})
    //登入表单提交
    beego.Router("/login", &controllers.LoginController{})
    //注册表单提交
    beego.Router("/register", &controllers.RegisterController{})
	//返回主页面
	beego.Router("/home.html",&controllers.HomeController{})
    //主页面查询
    beego.Router("/home",&controllers.HomeController{})
}
