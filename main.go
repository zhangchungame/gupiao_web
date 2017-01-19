package main

import (
	_ "gupiao_web/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gupiao_web/models"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/gupiao?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.Rimingxi))

}


func main() {
	beego.Run()
}

