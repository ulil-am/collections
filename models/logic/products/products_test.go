package products

import (
	"collections/structs"
	structsAPI "collections/structs/api/http"
	lStruct "collections/structs/logic"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertProduct_Success(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"
	reqCreateProducts := structsAPI.ReqCreateProducts{
		ProductCode: "01",
	}
	InsertProduct(ctxStruct, reqCreateProducts, &errCode)

	Convey("TestInsertProduct_Success", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
		})
	})
}

func TestInsertProduct_Error(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"
	reqCreateProducts := structsAPI.ReqCreateProducts{
		ProductCode: "productError",
	}
	InsertProduct(ctxStruct, reqCreateProducts, &errCode)

	Convey("TestInsertProduct_Error", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 1)
		})
	})
}

func TestGetAllProduct_Success(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	rows := GetAllProduct(&errCode)

	Convey("TestGetAllProduct_Success", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(len(rows), ShouldNotBeEmpty)
		})
	})
}

func TestGetProductList_Success(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	productCodes := make([]string, 0)
	productCodes = append(productCodes, "01")
	req := structsAPI.ReqInquiryProducts{
		productCodes,
	}
	res := GetProductList(req, &errCode)
	Convey("TestGetProductList_Success", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(len(res.Products), ShouldEqual, 1)
		})
	})
}

func TestGetProductList_Error(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	productCodes := make([]string, 0)
	productCodes = append(productCodes, "productError")
	req := structsAPI.ReqInquiryProducts{
		productCodes,
	}
	res := GetProductList(req, &errCode)
	Convey("TestGetProductList_Error", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 1)
			So(len(res.Products), ShouldEqual, 0)
		})
	})
}
