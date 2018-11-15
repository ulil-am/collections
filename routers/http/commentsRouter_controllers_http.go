package http

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["collections/controllers/http:CardsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:CardsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:CardsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:CardsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:GenerateController"] = append(beego.GlobalControllerRouter["collections/controllers/http:GenerateController"],
		beego.ControllerComments{
			Method: "LoanAppID",
			Router: `/loan-app-id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "GetProductList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["collections/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["collections/controllers/http:UserController"] = append(beego.GlobalControllerRouter["collections/controllers/http:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
