package http

import (
	"collections/helper"
	logicUser "collections/models/logic/user"
	"collections/structs"
	structsAPI "collections/structs/api/http"
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
}

// Post ...
// @Title Create
// @Description create User
// @Param	body		body 	structsAPI.ReqCreateUser		"body for User content"
// @Success 201 {object} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structsAPI.CreateUserInterface
		req          structsAPI.ReqCreateUser
		res          structsAPI.ResCreateUser
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
	logicUser.CreateUser(contextStruct, req, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	SendOutput(c.Ctx, res, errCode)

}
