package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gupiao_web/models"
	"fmt"
	"time"
)

type RimingxiController struct {
	beego.Controller
}



type resp struct {
	Code string
	Msg string
	Data interface{}
}
func (this *RimingxiController)GetDay()  {
	startDate:=this.GetString("startDate","20161201")
	endDate:=this.GetString("endDate","20161202")
	code:=this.GetString("code","000931")
	startTime,err:=time.Parse("20060102",startDate)
	if err!=nil{
		fmt.Println(err)
	}
	endTime,err:=time.Parse("20060102",endDate)
	if err!=nil{
		fmt.Println(err)
	}
	endTime=endTime.Add(3600*24-1)
	if len(code)!=6{
		fmt.Println("代码长度错误")
	}
	fmt.Println(startDate,endDate)
	db:=orm.NewOrm()
	infos:=make([]models.Rimingxi,0)
	db.QueryTable("rimingxi").Filter("code",code).Filter("date_int__gte",startTime.Unix()).Filter("date_int__lte",endTime.Unix()).OrderBy("date_int").All(&infos)
	resp:=resp{"200","success",infos}
	this.Data["json"]=&resp
	this.ServeJSON()
}
