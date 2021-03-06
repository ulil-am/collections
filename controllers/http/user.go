package http

import (
	"collections/helper"
	logicUser "collections/models/logic/user"
	"collections/structs"
	structAPI "collections/structs/api/http"
	"encoding/json"

	"github.com/astaxie/beego"
)

// UserController ...
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Create
// @Description create User
// @Param	body		body 	structAPI.ReqCreateUser		"body for User content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structAPI.CreateUserInterface
		req          structAPI.ReqCreateUser
		res          structAPI.ResCreateUser
	)

	rqBodyByte := helper.GetRqBodyRev(c.Ctx, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	err := json.Unmarshal(rqBodyByte, &reqInterface)
	beego.Debug("reqInterface ====>>>", reqInterface)
	if err != nil {
		structs.ErrorCode.UnexpectedError.String(&errCode)
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	reqInterface.ValidateRequest(&req, &errCode)
	beego.Debug("reqIface ====>>>", reqInterface)
	beego.Debug("req ====>>>", req)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	contextStruct := helper.ContextStruct(c.Ctx)
	res = logicUser.CreateUser(contextStruct, req, &errCode)
	beego.Debug("errAfterlogicCreateUser =======>>>>>>", errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	SendOutput(c.Ctx, res, errCode)

}

// Get ...
// @Title GetUser
// @Description get User by user_name
// @Param	user_name		path 	string	true		"The key for staticblock"
// @Success 200 {object} structAPI.ReqInquiryUser
// @Failure 403 :user_name is empty
// @router /:user_name [get]
func (c *UserController) Get() {
	errCode := make([]structs.TypeError, 0)

	var (
		req structAPI.ReqInquiryUser
		res structAPI.ResInquiryUser
	)

	req.UserName = c.Ctx.Input.Param(":user_name")
	contextStruct := helper.ContextStruct(c.Ctx)
	res = logicUser.GetUser(contextStruct, req, &errCode)

	SendOutput(c.Ctx, res, errCode)
}
