package http

import (
	"medium/helper"
	lModelGenAppID "medium/models/logic/genappid"
	"medium/structs"
	structsAPI "medium/structs/api/http"
	"encoding/json"

	"github.com/astaxie/beego"
)

// GenerateController operations for Generate
type GenerateController struct {
	beego.Controller
}

// URLMapping ...
func (c *GenerateController) URLMapping() {
	c.Mapping("LoanAppID", c.LoanAppID)
}

// LoanAppID ...
// @Title Create
// @Description create Generate
// @Param	body		body 	models.Generate	true		"body for Generate content"
// @Success 201 {object} models.Generate
// @Failure 403 body is empty
// @router /loan-app-id [post]
func (c *GenerateController) LoanAppID() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structsAPI.ReqLoanAppIDInterface
		req          structsAPI.ReqLoanAppID
		res          structsAPI.ResLoanAppID
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
	appID := lModelGenAppID.GenerateAppID(contextStruct, req, &errCode)
	res.AppID = appID

	SendOutput(c.Ctx, res, errCode)

}
