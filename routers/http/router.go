package http

import (
	ctrl "collections/controllers/http"

	"github.com/astaxie/beego"
)

func init() {
	Router()
}

// Router - Routing
func Router() {
	beego.InsertFilter("/*", beego.BeforeRouter, BeforeFunc, true)
	beego.ErrorHandler("404", pageNotFound)

	ns := beego.NewNamespace("/collections/v1",
		/*:STARTHTTP*/

		beego.NSRouter("/set_time",
			&ctrl.SetTimeController{},
			"post:SetTime"),
		beego.NSRouter("/set_spesific_time",
			&ctrl.SetTimeController{},
			"post:SetSpesificTime"),

		beego.NSNamespace("/generate",
			beego.NSInclude(
				&ctrl.GenerateController{},
			),
		),
		beego.NSNamespace("/products",
			beego.NSInclude(
				&ctrl.ProductsController{},
			),
		),
		beego.NSNamespace("/user/create",
			beego.NSInclude(
				&ctrl.UserController{},
			),
		),
		// With user_name

		beego.NSNamespace("/user/:user_name",
			beego.NSNamespace("/cards",
				beego.NSInclude(
					&ctrl.CardsController{},
				),
			),
			beego.NSNamespace("/cards/:card_number",
				beego.NSInclude(
					&ctrl.CardsController{},
				),
			),
		),

		// beego.NSNamespace("/cards",
		// 	beego.NSInclude(
		// 		&ctrl.CardsController{},
		// 	),
		// ),
		/*:ENDHTTP*/
	)

	beego.AddNamespace(ns)
	beego.SetStaticPath("/storages", "storages")
	beego.InsertFilter("/*", beego.FinishRouter, AfterFunc, true)
}
