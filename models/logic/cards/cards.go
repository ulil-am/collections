package cards

import (
	"collections/helper"
	"collections/structs"
	structAPI "collections/structs/api/http"
	structDb "collections/structs/db"
)

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
	err := DBCards.InsertCards(&doc)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc)
		helper.CheckErr(nmFunc+" ", err)
	}

	return
}
