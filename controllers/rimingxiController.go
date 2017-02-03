package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gupiao_web/models"
	"fmt"
	"time"
	"strconv"
)

type RimingxiController struct {
	beego.Controller
}



type resp struct {
	Code string
	Msg string
	Data interface{}
}

type Dayinfo struct {
	Time string
	Chengjiaojia float64
	Chengjiaoshou int64
}
func (this *RimingxiController)GetDay()  {
	startDate:=this.GetString("startDate","20161205")
	endDate:=this.GetString("endDate","20161205")
	code:=this.GetString("code","000007")
	startTime,err:=time.ParseInLocation("20060102",startDate, time.Local)
	if err!=nil{
		fmt.Println(err)
	}
	endTime,err:=time.ParseInLocation("20060102",endDate, time.Local)
	if err!=nil{
		fmt.Println(err)
	}
	endTime=endTime.AddDate(0,0,1)
	if len(code)!=6{
		fmt.Println("代码长度错误")
	}
	db:=orm.NewOrm()
	infos:=make([]models.Rimingxi,0)
	sql:="select * from rimingxi_"+code+" where date_int between "+strconv.FormatInt(startTime.Unix(),10)+" and "+strconv.FormatInt(endTime.Unix(),10)+" order by date_int asc"
	fmt.Println(sql)
	db.Raw(sql).QueryRows(&infos)
	result:=make([]Dayinfo,0)
	endTime=endTime.AddDate(0,0,-1)//恢复sql查询钱日期

	beginTime:=time.Unix(infos[0].Date_int,64)
	finishTime:=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),11,30,0,beginTime.Nanosecond(),time.Local)
	tmpDayinfo:=Dayinfo{}
	infoslen:=len(infos)
	index:=0
	fmt.Print(startTime.Unix())
	fmt.Print("-----")
	fmt.Println(endTime.Unix())
	for startTime.Unix()<=endTime.Unix(){
		beginTime=time.Date(startTime.Year(),startTime.Month(),startTime.Day(),9,30,0,startTime.Nanosecond(),time.Local)
		finishTime=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),11,30,0,beginTime.Nanosecond(),time.Local)
		for !beginTime.Equal(finishTime){
			tmpDayinfo=Dayinfo{}
			for i:=index;i<infoslen;i++{
				if(beginTime.Unix()>=infos[i].Date_int){
					tmpDayinfo.Chengjiaojia=infos[i].Chengjiaojia
					tmpDayinfo.Chengjiaoshou+=infos[i].Chengjiaoshou
				}else {
					break
				}
				index=i+1
			}
			tmpDayinfo.Time=beginTime.Format("02 15:04:05")
			beginTime=beginTime.Add(time.Minute)
			if tmpDayinfo.Chengjiaojia>0{

				result=append(result,tmpDayinfo)
			}
		}
		beginTime=time.Date(startTime.Year(),startTime.Month(),startTime.Day(),13,0,0,startTime.Nanosecond(),time.Local)
		finishTime=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),15,0,0,beginTime.Nanosecond(),time.Local)
		for !beginTime.Equal(finishTime){
			tmpDayinfo=Dayinfo{}
			for i:=index;i<infoslen;i++{
				if(beginTime.Unix()>=infos[i].Date_int){
					tmpDayinfo.Chengjiaojia=infos[i].Chengjiaojia
					tmpDayinfo.Chengjiaoshou+=infos[i].Chengjiaoshou
				}else {
					break
				}
				index=i+1
			}
			tmpDayinfo.Time=beginTime.Format("02 15:04:05")
			beginTime=beginTime.Add(time.Minute)
			if tmpDayinfo.Chengjiaojia>0{
				result=append(result,tmpDayinfo)
			}
		}
		startTime=startTime.AddDate(0,0,1)
		fmt.Print(startTime.Unix())
		fmt.Print("-----")
		fmt.Println(endTime.Unix())
	}





	//beginTime:=time.Unix(infos[0].Date_int,64)
	//beginTime=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),9,30,0,beginTime.Nanosecond(),time.Local)
	//fmt.Println(beginTime)
	//endTime=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),11,30,0,beginTime.Nanosecond(),time.Local)
	//fmt.Println(endTime)
	//tmpDayinfo:=Dayinfo{}
	//index:=0
	//infoslen:=len(infos)
	//for !beginTime.Equal(endTime){
	//	tmpDayinfo=Dayinfo{}
	//	for i:=index;i<infoslen;i++{
	//		if(beginTime.Unix()>=infos[i].Date_int){
	//			tmpDayinfo.Chengjiaojia=infos[i].Chengjiaojia
	//			tmpDayinfo.Chengjiaoshou+=infos[i].Chengjiaoshou
	//		}else {
	//			break
	//		}
	//		index=i+1
	//	}
	//	tmpDayinfo.Time=beginTime.Format("15:04:05")
	//	beginTime=beginTime.Add(time.Second)
	//	if tmpDayinfo.Chengjiaojia>0{
	//
	//		result=append(result,tmpDayinfo)
	//	}
	//}
	//beginTime=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),13,0,0,beginTime.Nanosecond(),time.Local)
	//fmt.Println(beginTime)
	//endTime=time.Date(beginTime.Year(),beginTime.Month(),beginTime.Day(),15,0,0,beginTime.Nanosecond(),time.Local)
	//fmt.Println(endTime)
	//for !beginTime.Equal(endTime){
	//	tmpDayinfo=Dayinfo{}
	//	for i:=index;i<infoslen;i++{
	//		if(beginTime.Unix()>=infos[i].Date_int){
	//			tmpDayinfo.Chengjiaojia=infos[i].Chengjiaojia
	//			tmpDayinfo.Chengjiaoshou+=infos[i].Chengjiaoshou
	//		}else {
	//			break
	//		}
	//		index=i+1
	//	}
	//	tmpDayinfo.Time=beginTime.Format("15:04:05")
	//	beginTime=beginTime.Add(time.Second)
	//	if tmpDayinfo.Chengjiaojia>0{
	//
	//		result=append(result,tmpDayinfo)
	//	}
	//}
	resp:=resp{"200","success",result}
	this.Data["json"]=&resp
	fmt.Println("req finish")
	this.ServeJSON()
}
