package routers

import (
	"gupiao_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.AutoRouter(&controllers.HtmlController{})
    beego.AutoRouter(&controllers.RimingxiController{})
}
