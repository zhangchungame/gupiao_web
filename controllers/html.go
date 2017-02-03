package controllers

import (
	"github.com/astaxie/beego"
)

type HtmlController struct {
	beego.Controller
}

func (this *HtmlController)Rikxian()  {
	this.Data["startDate"]=this.GetString("startDate","20161205")
	this.Data["endDate"]=this.GetString("endDate","20161205")
	this.Data["code"]=this.GetString("code","000007")
	this.TplName = "html/rikxian.html"
}