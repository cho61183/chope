package routers

import (
	"github.com/cho61183/chope_rpc/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
