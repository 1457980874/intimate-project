package controllers

import (
	"github.com/astaxie/beego"
	"intimate/DB"
	"intimate/models"
)

type LoginController struct {
	beego.Controller
}

func (u *LoginController) Get(){
	u.TplName="login.html"
}

func (u *LoginController) Post(){
	var user models.UserModels
	err:=u.ParseForm(&user)
	if err != nil {
		u.Ctx.WriteString("解析失败，请重试！")
		return
	}
	_, err = DB.QueryUser(user)
	if err != nil {
		u.Ctx.WriteString("登陆失败，请重试！")
		return
	}

}