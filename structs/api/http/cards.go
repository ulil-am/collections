package http

import (
	"collections/structs"

	"github.com/jannes-sa/customvalidator"
)

type (

	// ReqInquiryCards ...
	ReqInquiryCards struct {
		CardNumber []string `json:"card_number" example:"123456789012345"`
		UserName   string   `json:"user_name"`
	}

	// ResInquiryCards ...
	ResInquiryCards struct {
		UserName string          `json:"user_name"`
		Cards    []CardsResponse `json:"cards"`
	}

	//CardsResponse ...
	CardsResponse struct {
		CardNumber int    `json:"card_number"`
		Name       string `json:"name"`
		ExpiryDate string `json:"expiry_date"`
		Company    string `json:"company"`
	}

	//ReqCreateCards Struct for insert new card request
	ReqCreateCards struct {
		CardNumber int    `json:"card_number"`
		Name       string `json:"name"`
		ExpiryDate string `json:"expiry_date"`
		Company    string `json:"company"`
		UserName   string `json:"user_name"`
		// UserID     string `json:"user_id"`
	}

	//Cards Card struct
	Cards struct {
		CardNumber int    `json:"card_number"`
		Name       string `json:"name"`
		ExpiryDate string `json:"expiry_date"`
		Company    string `json:"company"`
		// UserName   string `json:"user_name"`
		UserID string `json:"user_id"`
	}

	// CreateCardsInterface ...
	CreateCardsInterface struct {
		CardNumber interface{} `json:"card_number" type:"int"`
		Name       interface{} `json:"name" type:"string"`
		ExpiryDate interface{} `json:"expiry_date" type:"string"`
		Company    interface{} `json:"company" type:"string"`
		UserName   interface{} `json:"user_name" type:"string"`
	}

	//RespCreateCards Struct for insert new card request
	RespCreateCards struct {
		UserName   string `json:"user_name"`
		CardNumber int    `json:"card_number"`
	}
)

// ValidateRequest ...
func (st *CreateCardsInterface) ValidateRequest(
	assignStruct *ReqCreateCards,
	errCode *[]structs.TypeError,
) {
	var paramsInterChange interface{} = *st
	codeError := customvalidator.Validate(
		paramsInterChange, assignStruct,
	)
	structs.GetCodeError(codeError, errCode)
}
