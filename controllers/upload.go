package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"os"
	"path"
	"time"
)

type UploadFireController struct {
	beego.Controller
}

func (f *UploadFireController) Get(){
	f.TplName="upload.html"
}

func (this *UploadFireController) Upload(){
	//获取文件
	f,h,_:=this.GetFile("myfile")
	fmt.Println(f)
	ex:=path.Ext(h.Filename)
	//验证文件的后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":true,
		".jpeg":true,
		"png":true,
	}
	if _,ok:=AllowExtMap[ex];!ok{
		this.Ctx.WriteString("后缀名不符合要求")
		return
	}
	//创建目录
	uploadDir:="static/img/"+time.Now().Format("2020/10/15")
	err:=os.MkdirAll(uploadDir,777)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("%v",err))
		return
	}
	rand.Seed(time.Now().UnixNano())
}