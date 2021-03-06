package http

import (
	"collections/routers/componenttest"
	"collections/structs"
	httpStructs "collections/structs/api/http"
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateSuccess(t *testing.T) {
	var req structs.ReqData

	var reqBody httpStructs.ReqLoanAppIDInterface
	reqBody.BranchCode = "0140"
	req.ReqBody = reqBody

	by, _ := json.Marshal(req)
	res := componenttest.SendHTTP(
		"POST",
		host+"/collections/v1/generate/loan-app-id",
		by,
	)

	Convey("TestGenerateSuccess", t, func() {
		Convey("Should Success", func() {
			So(len(res.Error), ShouldEqual, 0)
		})
	})
}
