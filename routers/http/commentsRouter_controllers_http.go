package http

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["medium/controllers/http:GenerateController"] = append(beego.GlobalControllerRouter["medium/controllers/http:GenerateController"],
		beego.ControllerComments{
			Method: "LoanAppID",
			Router: `/loan-app-id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["medium/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["medium/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["medium/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["medium/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "GetProductList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["medium/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["medium/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["medium/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["medium/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["medium/controllers/http:ProductsController"] = append(beego.GlobalControllerRouter["medium/controllers/http:ProductsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
