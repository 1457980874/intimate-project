package controllers

import (
	"github.com/astaxie/beego"
	"intimate/DB"
	"intimate/models"
)

type RegisterController struct {
	beego.Controller
}
//跳转页面
func (r *RegisterController) Get(){
	r.TplName="register.html"
}
//注册处理
func (r *RegisterController) Post(){
	//解析请求
	var user models.UserModels
	err:=r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("解析失败，请重试！")
		return
	}
	//保存用户信息
	_,err=DB.InsertUser(user)
	if err != nil {
		r.Ctx.WriteString("注册失败，请重试！")
		return
	}
	//加载登录页面模板
	r.TplName="login.html"
}
