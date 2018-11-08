package http

import (
	"collections/structs"

	"github.com/jannes-sa/customvalidator"
)

type (
	// ReqInquiryProducts ...
	ReqInquiryProducts struct {
		ProductCodes []string `json:"product"`
	}

	// ResInquiryProducts ...
	ResInquiryProducts struct {
		Products []Products `json:"product"`
	}

	// Products ...
	Products struct {
		ProductCode        string `json:"product_code"`
		ProductDescription string `json:"product_description"`
		MaximumLoanAmount  int    `json:"maximum_loan_amount"`
		MaximumTerm        int    `json:"maximum_term"`
		InterestRate       int    `json:"interest_rate"`
		CalculationMethod  string `json:"calculation_method"`
	}

	// ReqCreateProducts ...
	ReqCreateProducts struct {
		ProductCode        string `json:"product_code"`
		ProductDescription string `json:"product_description"`
		MaximumLoanAmount  int    `json:"maximum_loan_amount"`
		MaximumTerm        int    `json:"maximum_term"`
		InterestRate       int    `json:"interest_rate"`
		CalculationMethod  string `json:"calculation_method"`
	}

	// CreateProductsInterface ...
	CreateProductsInterface struct {
		ProductCode        interface{} `json:"product_code" type:"string"`
		ProductDescription interface{} `json:"product_description" type:"string"`
		MaximumLoanAmount  interface{} `json:"maximum_loan_amount" type:"int"`
		MaximumTerm        interface{} `json:"maximum_term" type:"int"`
		InterestRate       interface{} `json:"interest_rate" type:"int"`
		CalculationMethod  interface{} `json:"calculation_method" type:"string"`
	}
)

// ValidateRequest ...
func (st *CreateProductsInterface) ValidateRequest(
	assignStruct *ReqCreateProducts,
	errCode *[]structs.TypeError,
) {
	var paramsInterChange interface{} = *st
	codeError := customvalidator.Validate(
		paramsInterChange, assignStruct,
	)
	structs.GetCodeError(codeError, errCode)
}
