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
}

// Post ...
// @Title Create
// @Description Insert new cards
// @Param	body				body 	structAPI.ReqCreateCards	true	"Request Body of Create Cards"
// @Success 201 {object} models.Ecards
// @Failure 403 body is empty
// @Accept json
// @router / [post]
func (c *CardsController) Post() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structAPI.CreateCardsInterface
		req          structAPI.ReqCreateCards
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

	logicCards.InsertCards(req, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	SendOutput(c.Ctx, c.Data["json"], errCode)

}
