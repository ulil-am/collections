package cards

import (
	"collections/helper"
	"collections/structs"
	structAPI "collections/structs/api/http"
	structDb "collections/structs/db"
	structLogic "collections/structs/logic"
	"crypto/md5"
	"encoding/hex"
	"io"

	"errors"

	"github.com/astaxie/beego"
)

// CreateUser ...
func CreateUser(
	contextStruct structLogic.ContextStruct,
	reqCreateUser structAPI.ReqCreateUser,
	errCode *[]structs.TypeError,
) (resCreateUser structAPI.ResCreateUser) {
	var (
		doc    structDb.User
		nmFunc = "Create new user"
	)
	beego.Info(contextStruct.JobID, nmFunc)
	// beego.Debug("req =====>", reqCreateCards)

	// Check duplicate user ID

	// err2 := checkDuplicateUserName(reqCreateUser, errCode)
	// if err2 != nil {
	// 	structs.ErrorCode.UserNameAlreadyExist.String(errCode)
	// 	return
	// }
	counter := getInc(reqCreateUser, errCode)

	beego.Debug("errAftergetInc =======>>>>>>", errCode)

	hash := md5.New()
	io.WriteString(hash, reqCreateUser.Password)

	hashinBytes := hash.Sum(nil)[:16]
	maskedPwd := hex.EncodeToString(hashinBytes)

	doc.UserID = counter
	doc.UserName = reqCreateUser.UserName
	doc.Password = maskedPwd
	doc.Email = reqCreateUser.Email

	beego.Debug("doc ======> ", doc)

	err := DBUser.CreateUser(doc)
	beego.Debug("err after insert ========> ", err)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, nmFunc, logicName)
	}

	resCreateUser.UserID = doc.UserID
	resCreateUser.UserName = doc.UserName
	return
}

// GetUser ...
func GetUser(
	contextStruct structLogic.ContextStruct,
	reqInqUser structAPI.ReqInquiryUser,
	errCode *[]structs.TypeError,
) (resInqUser structAPI.ResInquiryUser) {
	var (
		doc      string
		nmFunc   = "Get user details"
		dataUser structDb.User
		err      error
	)
	beego.Info(contextStruct.JobID, nmFunc)

	doc = reqInqUser.UserName

	beego.Debug("doc ======> ", doc)

	dataUser, err = DBUser.GetUserByUserName(doc)
	beego.Debug("err after inq ========> ", err)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, nmFunc, logicName)
	}

	resInqUser.UserName = dataUser.UserName
	resInqUser.UserID = dataUser.UserID
	resInqUser.Email = dataUser.Email

	return
}

func getInc(
	reqCreateUser structAPI.ReqCreateUser,
	errCode *[]structs.TypeError,
) (counterID int) {
	var row structDb.CounterUser
	row.Counter = 0
	row.UserCounter = "card_user"
	saveToCounterIfNotExist(row)

	counter, err := DBCounter.GetInc(row)
	beego.Debug("err after GetInc to DB ========> ", err)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error())
	}
	counterID = counter
	beego.Debug("counterID ====>", counterID)
	return
}

func saveToCounterIfNotExist(
	row structDb.CounterUser,
) {
	err := DBCounter.InsertCounter(row)
	helper.CheckErr("data already exist", err)
}

func checkDuplicateUserName(
	reqCreateUser structAPI.ReqCreateUser,
	errCode *[]structs.TypeError,
) (err error) {
	var (
		row  structDb.User
		err2 error
	)
	row, err2 = DBUser.GetUserByUserName(reqCreateUser.UserName)

	if err2 == nil {
		if row.UserName == reqCreateUser.UserName {
			err = errors.New("Username Already Exist")
		}
	}

	return
}
