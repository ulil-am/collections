package http

import (
	"collections/structs"

	"github.com/jannes-sa/customvalidator"
)

type (

	// ReqCreateUser Struct for create new user request
	ReqCreateUser struct {
		UserName string `json:"user_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// ResCreateUser ...
	ResCreateUser struct {
		UserName string `json:"user_name"`
		UserID   int    `json:"user_id"`
	}

	// User ...
	User struct {
		UserName string `json:"user_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		UserID   int    `json:"user_id"`
	}

	// CreateUserInterface ...
	CreateUserInterface struct {
		UserName interface{} `json:"user_name" type:"string"`
		Email    interface{} `json:"email" type:"string"`
		Password interface{} `json:"password" type:"string"`
		// UserID int `json:"user_id"`
	}
)

// ValidateRequest ...
func (st *CreateUserInterface) ValidateRequest(
	assignStruct *ReqCreateUser,
	errCode *[]structs.TypeError,
) {
	var paramsInterChange interface{} = *st
	codeError := customvalidator.Validate(
		paramsInterChange, assignStruct,
	)
	structs.GetCodeError(codeError, errCode)
}
