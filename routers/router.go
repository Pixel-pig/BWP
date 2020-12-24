package routers

import (
	"BTCproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//默认
	beego.Router("/", &controllers.MainController{})
	//注册
	beego.Router("/index", &controllers.Index{})
	//登入
	beego.Router("/login",&controllers.Login{})
	//注册页面
	beego.Router("/index.html",&controllers.Index{})
	//方法页面
	beego.Router("/Btc_home.html",&controllers.Btc_home{})
	//功能页面
	beego.Router("/obtain.html",&controllers.Obtain{})
	//调用方法
	beego.Router("/obtain",&controllers.Obtain{})
}
