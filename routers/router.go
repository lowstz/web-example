package routers

import (
	"github.com/lowstz/web-example/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
