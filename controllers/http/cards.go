package http

import (
	"collections/helper"
	logicCards "collections/models/logic/cards"
	"collections/structs"
	structAPI "collections/structs/api/http"
	"encoding/json"

	"github.com/astaxie/beego"
)

// CardsController controller operations for cards
type CardsController struct {
	beego.Controller
}

// URLMapping ...
func (c *CardsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Get ...
// @Title GetCard
// @Description get Cards by card_number
// @Param	card_number				query	int	1	"Card Number"	"?card_number=1"
// @Success 200 {object} structAPI.ReqInquiryCards
// @Failure 403 :card_number is empty
// @router / [get]
func (c *CardsController) Get() {
	errCode := make([]structs.TypeError, 0)

	var (
		req structAPI.ReqInquiryCards
		res structAPI.ResInquiryCards
	)

	req.UserName = c.Ctx.Input.Param(":user_name")

	req.CardNumber = c.Ctx.Request.URL.Query()["card_number"][0]
	contextStruct := helper.ContextStruct(c.Ctx)
	res = logicCards.GetCardInfo(contextStruct, req, &errCode)

	SendOutput(c.Ctx, res, errCode)
}

// Post ...
// @Title Create
// @Description Insert new cards
// @Param	body				body 	structAPI.ReqCreateCards	true	"Request Body of Create Cards"
// @Success 201 {object} models.Cards
// @Failure 403 body is empty
// @Accept json
// @router / [post]
func (c *CardsController) Post() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structAPI.CreateCardsInterface
		req          structAPI.ReqCreateCards
		res          structAPI.RespCreateCards
	)
	rqBodyByte := helper.GetRqBodyRev(c.Ctx, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}
	reqInterface.UserName = c.Ctx.Input.Param(":user_name")
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
	res = logicCards.InsertCards(contextStruct, req, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	SendOutput(c.Ctx, res, errCode)

}
