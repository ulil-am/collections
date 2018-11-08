package http

import (
	"collections/structs"

	"github.com/jannes-sa/customvalidator"
)

type (
	// ReqLoanAppIDInterface - ReqLoanAppIDInterface Struct
	ReqLoanAppIDInterface struct {
		BranchCode interface{} `json:"branch_code" type:"string" validate:"required=0301,stringnumericonly=0302,len=4=0303"`
	}
	// ReqLoanAppID - ReqLoanAppID
	ReqLoanAppID struct {
		BranchCode string
	}

	// ResLoanAppID - ResLoanAppID
	ResLoanAppID struct {
		AppID string `json:"application_id"`
	}
)

// ValidateRequest ...
func (st *ReqLoanAppIDInterface) ValidateRequest(
	assignStruct *ReqLoanAppID,
	errCode *[]structs.TypeError,
) {
	var paramsInterChange interface{} = *st
	codeError := customvalidator.Validate(
		paramsInterChange, assignStruct,
	)
	structs.GetCodeError(codeError, errCode)
}
