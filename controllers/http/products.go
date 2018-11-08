package http

import (
	"medium/helper"
	logicProduct "medium/models/logic/products"
	"medium/structs"
	structsAPI "medium/structs/api/http"
	"encoding/json"
)

import (
	"github.com/astaxie/beego"
)

// ProductsController operations for Products
type ProductsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetProductList", c.GetProductList)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Products
// @Param	body		body 	models.Products	true		"body for Products content"
// @Success 201 {object} models.Products
// @Failure 403 body is empty
// @router / [post]
func (c *ProductsController) Post() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structsAPI.CreateProductsInterface
		req          structsAPI.ReqCreateProducts
	)

	rqBodyByte := helper.GetRqBodyRev(c.Ctx, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	err := json.Unmarshal(rqBodyByte, &reqInterface)
	if err != nil {
		structs.ErrorCode.UnexpectedError.String(&errCode)
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	reqInterface.ValidateRequest(&req, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	contextStruct := helper.ContextStruct(c.Ctx)
	logicProduct.InsertProduct(contextStruct, req, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	SendOutput(c.Ctx, c.Data["json"], errCode)

}

// GetOne ...
// @Title GetOne
// @Description get Products by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Products
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductsController) GetOne() {
	errCode := make([]structs.TypeError, 0)

	id := c.Ctx.Input.Param(":id")
	beego.Debug(id)

	SendOutput(c.Ctx, c.Data["json"], errCode)

}

// GetProductList ...
// @Title GetProductList
// @Description get Products
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Products
// @Failure 403
// @router / [get]
func (c *ProductsController) GetProductList() {
	errCode := make([]structs.TypeError, 0)

	var (
		req structsAPI.ReqInquiryProducts
		res structsAPI.ResInquiryProducts
	)

	req.ProductCodes = c.Ctx.Request.URL.Query()["product"]
	res = logicProduct.GetProductList(req, &errCode)

	SendOutput(c.Ctx, res, errCode)
}

// Put ...
// @Title Put
// @Description update the Products
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Products	true		"body for Products content"
// @Success 200 {object} models.Products
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Products
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductsController) Delete() {

}
