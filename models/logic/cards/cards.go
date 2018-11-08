package cards

import (
	"collections/helper"
	"collections/structs"
	structAPI "collections/structs/api/http"
	structDb "collections/structs/db"
)

// InsertCards ...
func InsertCards(
	reqCreateCards structAPI.ReqCreateCards,
	errCode *[]structs.TypeError,
) {
	var (
		doc    structDb.Cards
		nmFunc = "Insert new cards"
	)

	doc.CardNumber = reqCreateCards.CardNumber
	doc.Company = reqCreateCards.Company
	doc.ExpiryDate = reqCreateCards.ExpiryDate
	doc.Name = reqCreateCards.Name
	err := DBCards.InsertCards(doc)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc)
		helper.CheckErr(nmFunc+" ", err)
	}

	return
}

// GetCardInfo ...
func GetCardInfo(reqInquiryCards structAPI.ReqInquiryCards,
	errCode *[]structs.TypeError,
) (resInquiryCards structAPI.ResInquiryCards) {
	var (
		nmFunc = "GetCardInfo"
		rows   []structDb.Cards
		err    error
	)

	if len(reqInquiryCards.CardNumber) == 0 {
		structs.ErrorCode.MissingField.String(errCode, nmFunc)
		return
	} else {
		rows, err = DBCards.GetCardByCardNumber(reqInquiryCards.CardNumber)
		if err != nil {
			structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
			return
		}
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
