package controllers

import (
	"github.com/astaxie/beego"
)

type HtmlController struct {
	beego.Controller
}

func (this *HtmlController)Rikxian()  {
	this.Data["content"] = "value"
	this.Data["content2"] = "aaa"
	this.TplName = "html/rikxian.html"
}