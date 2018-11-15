package cards

import (
	"collections/structs"
	structAPI "collections/structs/api/http"
	structLogic "collections/structs/logic"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertCards_Success(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct structLogic.ContextStruct
	ctxStruct.JobID = "jobid1"
	reqInsertCards := structAPI.ReqCreateCards{
		CardNumber: 1311,
	}
	InsertCards(ctxStruct, reqInsertCards, &errCode)

	Convey("TestInsertCards_Success", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
		})
	})
}

func TestInsertCards_Error(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct structLogic.ContextStruct
	ctxStruct.JobID = "jobid2"
	reqInsertCards := structAPI.ReqCreateCards{
		CardNumber: 0,
	}
	InsertCards(ctxStruct, reqInsertCards, &errCode)

	Convey("TestInsertCards_Eerror", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 1)
		})
	})
}

func TestGetCardInfo_Success(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct structLogic.ContextStruct
	ctxStruct.JobID = "jobid3"
	reqInquiryCards := structAPI.ReqInquiryCards{
		CardNumber: "1",
	}
	res := GetCardInfo(ctxStruct, reqInquiryCards, &errCode)

	Convey("TestGetCardInfo_Success", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(len(res.Cards), ShouldEqual, 1)
		})
	})
}

func TestGetCardInfo_Error(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct structLogic.ContextStruct
	ctxStruct.JobID = "jobid4"
	reqInquiryCards := structAPI.ReqInquiryCards{
		CardNumber: "0",
	}
	res := GetCardInfo(ctxStruct, reqInquiryCards, &errCode)

	Convey("TestGetCardInfo_Success", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 1)
			So(len(res.Cards), ShouldEqual, 0)
		})
	})
}
