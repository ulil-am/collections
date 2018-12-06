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
	beego.Debug("req =====>", reqCreateCards)
	if reqCreateCards.CardNumber == 0 {
		structs.ErrorCode.CardNumberZero.String(errCode)
	} else {

		// Get User ID

		userID := getUserID(reqCreateCards.UserName, errCode)

		doc.CardNumber = reqCreateCards.CardNumber
		doc.Company = reqCreateCards.Company
		doc.ExpiryDate = reqCreateCards.ExpiryDate
		doc.Name = reqCreateCards.Name
		doc.UserID = userID
		beego.Debug("doc =====>", doc)
		err := DBCards.InsertCards(doc)
		if err != nil {
			structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc)
			helper.CheckErr(nmFunc+" ", err)
		}
		resInsertCard.CardNumber = doc.CardNumber
		resInsertCard.UserName = reqCreateCards.UserName
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

	// get userID
	userID := getUserID(reqInquiryCards.UserName, errCode)
	if len(reqInquiryCards.CardNumber) == 0 {
		rows = GetAllCards(userID, errCode)
		if len(*errCode) != 0 {
			return
		}
		if len(rows) == 0 {
			structs.ErrorCode.CardNotExist.String(errCode, "Get All Cards", logicName)
		}
	} else {
		cardNumber, errConv := strconv.Atoi(reqInquiryCards.CardNumber[0])
		if errConv != nil {
			structs.ErrorCode.MissingField.String(errCode, errConv.Error())
			return
		}
		rows, err = DBCards.GetCardByCardNumber(cardNumber, userID)
		if err != nil {
			structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
			return
		}

		if len(rows) == 0 {
			structs.ErrorCode.CardNotExist.String(errCode, nmFunc, logicName)
			return
		}
	}

	resInquiryCards.UserName = reqInquiryCards.UserName
	for _, val := range rows {
		cards := structAPI.CardsResponse{
			CardNumber: val.CardNumber,
			ExpiryDate: val.ExpiryDate,
			Company:    val.Company,
			Name:       val.Name,
		}
		resInquiryCards.Cards = append(resInquiryCards.Cards, cards)
	}

	return

}

// GetAllCards ...
func GetAllCards(u int, errCode *[]structs.TypeError) (rows []structDb.Cards) {
	rows, err := DBCards.GetAllCards(u)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), "Get All Cards", logicName)
	}

	return
}

func getUserID(
	un string,
	errCode *[]structs.TypeError,
) (userID int) {
	var (
		row  structDb.User
		err2 error
	)
	row, err2 = DBUser.GetUserByUserName(un)

	if err2 != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err2.Error())
	}

	userID = row.UserID

	return
}
