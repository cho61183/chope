// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/cho61183/chope_rpc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/rpc",

		beego.NSNamespace("/user_info",
			beego.NSInclude(
				&controllers.UserInfoController{},
			),
		),

		beego.NSNamespace("/restaurant_info",
			beego.NSInclude(
				&controllers.RestaurantInfoController{},
			),
		),

		beego.NSNamespace("/restaurant_cms",
			beego.NSInclude(
				&controllers.RestaurantCmsController{},
			),
		),

		beego.NSNamespace("/points_log",
			beego.NSInclude(
				&controllers.PointsLogController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
