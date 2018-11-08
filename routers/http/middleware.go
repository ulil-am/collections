package http

import (
	ctrl "collections/controllers/http"
	"collections/helper"
	"collections/helper/constant"
	"collections/helper/timetn"
	"collections/structs"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

// Middleware - Middleware Struct
type Middleware struct{}

func (m *Middleware) log(c *context.Context) {
	jobID := helper.GetJobID(c)
	go func() {
		beego.Info(
			"REQ_JOBID", jobID,
			"REQ_URL", jobID, c.Input.URL(),
			"REQ_HEADER", jobID, helper.HeaderAll(c),
			"REQ_BODY",
			jobID,
			strings.Replace(
				string(c.Input.RequestBody),
				"\n", "", -1),
		)
	}()
}

func (m *Middleware) initialHeader(c *context.Context) {
	m.handleRoundTrip(c)
	m.handleReqID(c)
	m.handleJobID(c)
	m.setMessageID(c)
}

func (m *Middleware) handleRoundTrip(c *context.Context) {
	ms := timetn.Now().UnixNano() / int64(time.Millisecond)
	c.Input.SetData("x-roundtrip", strconv.FormatInt(ms, 10))
}

func (m *Middleware) handleReqID(c *context.Context) {
	reqID := c.Input.Header("x-request-id")
	c.Output.Header("x-request-id", reqID)
}

func (m *Middleware) handleJobID(c *context.Context) {
	jobID := c.Input.Header("x-job-id")
	newRequest := 0
	if jobID == "" {
		newRequest = 1
		c.Input.SetData("job-id", helper.GenJobID())
		c.Input.SetData("new_request", newRequest)
	} else {
		c.Input.SetData("job-id", jobID)
		c.Input.SetData("new_request", newRequest)
	}
}

func (m *Middleware) setMessageID(c *context.Context) {
	uuID := helper.GetUUID()
	c.Input.SetData("message-id", uuID)
}

func (m *Middleware) authToken(c *context.Context) {
	if constant.AUTH == "" || constant.AUTH == "0" {
		return
	}

	if c.Input.URL() == "/token" {
		generateToken(c)
		return
	}

	errCode := make([]structs.TypeError, 0)
	tokenString := c.Input.Header("token")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(`jwtkey`), nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			beego.Warning("That's not even a token", tokenString)
			structs.ErrorCode.TokenGenerateDenied.String(&errCode)
			ctrl.SendOutput(c, nil, errCode)
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			beego.Warning("Token Already Expired or Not Valid Yet", tokenString)
			structs.ErrorCode.TokenExpired.String(&errCode)
			ctrl.SendOutput(c, nil, errCode)
			return
		} else {
			beego.Warning("Couldn't handle this token 1", err, tokenString)
			structs.ErrorCode.TokenInvalid.String(&errCode)
			ctrl.SendOutput(c, nil, errCode)
			return
		}
	}
}

func (m *Middleware) checkRqHeader(c *context.Context) {
	if c.Input.GetData("ignoreheader") == nil {

		headerAll := helper.HeaderAll(c)
		var reqHeaderStruct structs.ReqHTTPHeader
		err := json.Unmarshal([]byte(headerAll), &reqHeaderStruct)
		if err != nil {
			log.Println(err)
		}

		errVal := reqHeaderStruct.ValidateReqHTTPHeader()
		if errVal != nil {
			var errCode []structs.TypeError
			structs.ErrorCode.RequestMalformed.String(&errCode)

			filErrCode := make([]structs.FilterErrorCode, 0)
			if len(errCode) > 0 {
				filErrCode = filterErrorCode(errCode)
			}

			var respData structs.RespData
			respData.Error = filErrCode

			respByte, err := json.Marshal(respData)
			if err != nil {
				log.Println(err)
			}

			// resHeader := helper.ConstructHTTPHeader(c)
			// resHeaderByte, err := json.Marshal(resHeader)
			// if err != nil {
			// 	log.Println(err)
			// }

			c.Output.SetStatus(400)
			err = c.Output.Body(respByte)
			if err != nil {
				log.Println(err)
			}
		}

	}
}

func generateToken(c *context.Context) {
	errCode := make([]structs.TypeError, 0)

	type tokenClaim struct {
		Data string `json:"data"`
		jwt.StandardClaims
	}

	timeExp, err := strconv.Atoi(constant.AUTHEXP)
	helper.CheckErr("failed strconv AUTHEXP middleware http", err)
	if err != nil {
		structs.ErrorCode.TokenGenerateFailed.String(&errCode)
		ctrl.SendOutput(c, nil, errCode)
		return
	}

	// Create the Claims
	claims := tokenClaim{
		"customDataHere",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(timeExp) * time.Minute).Unix(),
			Issuer:    "jannes",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(`jwtkey`))
	if err != nil {
		beego.Warning("failed generate token with signkey")
		structs.ErrorCode.TokenGenerateFailed.String(&errCode)
		ctrl.SendOutput(c, nil, errCode)
		return
	}

	var respToken struct {
		Token string `json:"token"`
	}
	respToken.Token = tokenString
	ctrl.SendOutput(c, respToken, errCode)

	return
}

// BeforeFunc - Execute before all router call
func BeforeFunc(c *context.Context) {
	var m Middleware

	m.initialHeader(c)
	m.log(c)
	m.authToken(c)
	m.checkRqHeader(c)
}

// AfterFunc to execute progress after response
func AfterFunc(c *context.Context) {

}

// pageNotFound ..
func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	_, err := rw.Write([]byte(""))
	if err != nil {
		beego.Warning("NOT FOUND ERROR")
	}
}

func filterErrorCode(arrErrType []structs.TypeError) []structs.FilterErrorCode {
	var filter []structs.FilterErrorCode
	for _, val := range arrErrType {
		filter = append(filter, structs.FilterErrorCode{
			Code:    val.Code,
			Message: val.Message,
		})
	}
	return filter
}
