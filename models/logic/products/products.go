package products

import (
	"medium/helper"
	"medium/structs"
	structsAPI "medium/structs/api/http"
	dbStruct "medium/structs/db"
	lStruct "medium/structs/logic"
)

// GetAllProduct - GetAllProduct
func GetAllProduct(errCode *[]structs.TypeError) (rows []dbStruct.Product) {
	rows, err := DBProducts.GetAllProduct()
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), "GetAllProduct ", logicName)
	}

	return
}

// GetProductList - Get Product base on product code
func GetProductList(reqInquiryProducts structsAPI.ReqInquiryProducts, errCode *[]structs.TypeError,
) (resInquiryProducts structsAPI.ResInquiryProducts) {
	var (
		nmFunc = "GetProductList"
	)

	rows, err := DBProducts.GetProductByProductCodes(reqInquiryProducts.ProductCodes)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
	}

	for _, val := range rows {
		product := structsAPI.Products{
			ProductCode:        val.CampaignCode,
			ProductDescription: val.CampaignDesc,
			CalculationMethod:  val.InstallmentCalculationMethod,
			InterestRate:       val.InterestRate,
			MaximumLoanAmount:  val.MaxLoanAmount,
			MaximumTerm:        val.MaxTerm,
		}
		resInquiryProducts.Products = append(resInquiryProducts.Products, product)
	}

	return
}

// InsertProduct - InsertProduct
func InsertProduct(
	contextStruct lStruct.ContextStruct,
	reqCreateProducts structsAPI.ReqCreateProducts,
	errCode *[]structs.TypeError,
) {
	var (
		nmFunc = "InsertProduct"
		row    dbStruct.Product
	)

	row.CampaignCode = reqCreateProducts.ProductCode
	row.CampaignDesc = reqCreateProducts.ProductDescription
	row.InstallmentCalculationMethod = reqCreateProducts.CalculationMethod
	row.InterestRate = reqCreateProducts.InterestRate
	row.MaxLoanAmount = reqCreateProducts.MaximumLoanAmount
	row.MaxTerm = reqCreateProducts.MaximumTerm

	err := DBProducts.InsertProduct(&row)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc, logicName)
		helper.CheckErr(nmFunc+" "+logicName, err)

		return
	}
}
