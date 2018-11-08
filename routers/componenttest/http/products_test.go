package http

import (
	"collections/helper/constant"
	"collections/routers/componenttest"
	"collections/structs"
	httpStructs "collections/structs/api/http"
	"encoding/json" // "strconv"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	// "github.com/astaxie/beego"
)

func TestCreateProductsSuccess(t *testing.T) {

	var req structs.ReqData
	var reqBody httpStructs.ReqCreateProducts
	reqBody.ProductCode = "01"
	reqBody.ProductDescription = "desc"
	reqBody.InterestRate = 1
	reqBody.MaximumLoanAmount = 1
	reqBody.MaximumTerm = 1
	reqBody.CalculationMethod = "a"
	req.ReqBody = reqBody
	by, _ := json.Marshal(req)

	res := componenttest.SendHTTP(
		"POST",
		host+"/"+constant.DOMAINNAME+"/v1/products",
		by,
	)

	Convey("TestCreateProductsSuccess", t, func() {
		Convey("Should Success", func() {
			So(len(res.Error), ShouldEqual, 0)
		})
	})
}

func TestGetProductList_Success(t *testing.T) {

	var by []byte
	productCode := "01"
	res := componenttest.SendHTTP(
		"GET",
		host+"/"+constant.DOMAINNAME+"/v1/products?product="+productCode,
		by,
	)
	var resGetProductList httpStructs.ResInquiryProducts
	resMarshal, _ := json.Marshal(res.ResponseBody)
	json.Unmarshal(resMarshal, &resGetProductList)
	beego.Warning(resGetProductList)
	Convey("TestCreateProductsSuccess", t, func() {
		Convey("Should Success", func() {
			So(len(res.Error), ShouldEqual, 0)
			So(resGetProductList.Products[0].ProductCode, ShouldEqual, productCode)
		})
	})
}

func TestGetProductList_Error(t *testing.T) {

	var by []byte
	productCode := "productError"
	res := componenttest.SendHTTP(
		"GET",
		host+"/"+constant.DOMAINNAME+"/v1/products?product="+productCode,
		by,
	)
	var resGetProductList httpStructs.ResInquiryProducts
	resMarshal, _ := json.Marshal(res.ResponseBody)
	json.Unmarshal(resMarshal, &resGetProductList)
	Convey("TestCreateProductsSuccess", t, func() {
		Convey("Should Success", func() {
			So(len(res.Error), ShouldEqual, 1)
			So(len(resGetProductList.Products), ShouldBeZeroValue)
		})
	})
}
