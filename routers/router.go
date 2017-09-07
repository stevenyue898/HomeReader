package routers

import (
	"homereader/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/getDir", &controllers.GetDirController{})
}
