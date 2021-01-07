package routers

import (
	"BWPjy/controllers"
	"github.com/astaxie/beego"
)
func init() {
    beego.Router("/", &controllers.MainController{})
    //处理登录页面请求
    beego.Router("/login",&controllers.LonginControllers{})
    //处理注册页面请求
    beego.Router("/regist",&controllers.RegistControllers{})
    //home页面请求
    beego.Router("/home",&controllers.HomeControllers{})
}
