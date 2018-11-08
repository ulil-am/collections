package genappid

import (
	"medium/structs"
	apiStruct "medium/structs/api/http"
	lStruct "medium/structs/logic"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInc(t *testing.T) {
	var req apiStruct.ReqLoanAppID
	errCode := make([]structs.TypeError, 0)

	req.BranchCode = "0140"

	getInc(req, &errCode)

	Convey("TestGetInc", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
		})
	})
}

func TestGenerateAppID(t *testing.T) {
	var contextStruct lStruct.ContextStruct
	var req apiStruct.ReqLoanAppID
	errCode := make([]structs.TypeError, 0)

	req.BranchCode = "0140"
	appID := GenerateAppID(contextStruct, req, &errCode)

	Convey("TestGenerateAppID", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(appID, ShouldEqual, "DL014010180001")
		})
	})
}
