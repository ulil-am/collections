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
	counter := getInc(reqCreateUser, errCode)

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
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, nmFunc, logicName)
	}

	resCreateUser.UserName = doc.UserName
	return
}

func getInc(
	reqCreateUser structAPI.ReqCreateUser,
	errCode *[]structs.TypeError,
) (counterID int) {
	var row structDb.CounterUser
	row.Counter = 0
	saveToCounterIfNotExist(row)

	counter, err := DBCounter.GetInc(row)
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
