package genappid

import (
	"medium/helper"
	"medium/helper/timetn"
	"medium/structs"
	apiStruct "medium/structs/api/http"
	dbStruct "medium/structs/db"
	lStruct "medium/structs/logic"
	"strconv"

	"github.com/astaxie/beego"
)

var (
	digitsCounter = 4
)

// GenerateAppID - Generate Application ID
func GenerateAppID(
	contextStruct lStruct.ContextStruct,
	req apiStruct.ReqLoanAppID,
	errCode *[]structs.TypeError,
) (appID string) {
	beego.Info(contextStruct.JobID)
	counter := getInc(req, errCode)

	month := strconv.Itoa(int(timetn.Now().Month()))
	year := strconv.Itoa(timetn.Now().Year())[2:]

	appID = "DL" + req.BranchCode + month + year + counter

	return
}

func getInc(
	req apiStruct.ReqLoanAppID,
	errCode *[]structs.TypeError,
) (counterStr string) {
	var row dbStruct.Counter
	row.BranchCode = req.BranchCode
	row.Month = int(timetn.Now().Month())
	row.Counter = 1

	saveToCounterIfNotExist(row)

	counter, err := DBGenappid.GetInc(row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, "failed increment @getInc",
			logicName)
		return
	}

	convCounter := strconv.Itoa(counter)
	for i := 0; i < (digitsCounter - len(convCounter)); i++ {
		counterStr += "0"
	}
	counterStr += convCounter

	return
}

func saveToCounterIfNotExist(
	row dbStruct.Counter,
) {
	err := DBGenappid.InsertServiceLog(row)
	helper.CheckErr("data already exists", err)
}
