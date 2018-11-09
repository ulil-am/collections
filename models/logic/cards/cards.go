package cards

import (
	"collections/helper"
	"collections/structs"
	structAPI "collections/structs/api/http"
	structDb "collections/structs/db"
	structLogic "collections/structs/logic"
	"strconv"

	"github.com/astaxie/beego"
)

// InsertCards ...
func InsertCards(
	contextStruct structLogic.ContextStruct,
	reqCreateCards structAPI.ReqCreateCards,
	errCode *[]structs.TypeError,
) (resInsertCard structAPI.RespCreateCards) {
	var (
		doc    structDb.Cards
		nmFunc = "Insert new cards"
	)
	beego.Info(contextStruct.JobID, nmFunc)

	if reqCreateCards.CardNumber == 0 {
		structs.ErrorCode.CardNumberZero.String(errCode)
	} else {
		doc.CardNumber = reqCreateCards.CardNumber
		doc.Company = reqCreateCards.Company
		doc.ExpiryDate = reqCreateCards.ExpiryDate
		doc.Name = reqCreateCards.Name
		err := DBCards.InsertCards(doc)
		if err != nil {
			structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc)
			helper.CheckErr(nmFunc+" ", err)
		}
		resInsertCard.CardNumber = doc.CardNumber
		return
	}

	return
}

// GetCardInfo ...
func GetCardInfo(
	contextStruct structLogic.ContextStruct,
	reqInquiryCards structAPI.ReqInquiryCards,
	errCode *[]structs.TypeError,
) (resInquiryCards structAPI.ResInquiryCards) {
	var (
		nmFunc = "GetCardInfo"
		rows   []structDb.Cards
		err    error
	)
	beego.Info(contextStruct.JobID, nmFunc)

	cardNumber, errConv := strconv.Atoi(reqInquiryCards.CardNumber)
	if errConv != nil {
		structs.ErrorCode.MissingField.String(errCode, errConv.Error())
		return
	}

	if cardNumber == 0 {
		structs.ErrorCode.CardNumberZero.String(errCode, nmFunc)
		return
	}

	rows, err = DBCards.GetCardByCardNumber(cardNumber)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		return
	}

	if len(rows) == 0 {
		structs.ErrorCode.CardNotExist.String(errCode, nmFunc, logicName)
		return
	}

	for _, val := range rows {
		cards := structAPI.Cards{
			CardNumber: val.CardNumber,
			ExpiryDate: val.ExpiryDate,
			Company:    val.Company,
			Name:       val.Name,
		}
		resInquiryCards.Cards = append(resInquiryCards.Cards, cards)
	}

	return

}
