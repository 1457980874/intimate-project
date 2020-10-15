package routers

import (
	"intimate/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //注册接口请求
    beego.Router("/register",&controllers.RegisterController{})
    //登录接口请求
    beego.Router("/login",&controllers.LoginController{})
}
