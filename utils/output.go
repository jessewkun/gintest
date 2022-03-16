package utils

import (
	"encoding/json"
	"fmt"
	"gintest/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) *APIException

type APIException struct {
	HttpStatus int         `json:"-"`
	RawErr     error       `json:"-"`
	Tag        string      `json:"-"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (ae APIException) String() string {
	jsonBytes, err := json.Marshal(ae)
	if err != nil {
		return fmt.Sprintf("HttpStatus: %d, RawErr: %s, Tag: %s, Code: %d, Message: %s, Data: %s", ae.HttpStatus, ae.RawErr, ae.Tag, ae.Code, ae.Message, ae.Data)
	}
	return string(jsonBytes)
}

func NewAPIException(code int, err error) *APIException {
	message, ok := config.ERROR_MESSAGE_MAP[code]
	if !ok {
		code = config.ERROR_SYSTEM
		message = config.ERROR_MESSAGE_MAP[config.ERROR_SYSTEM]
	}
	return &APIException{
		HttpStatus: http.StatusOK,
		RawErr:     err,
		Tag:        "APIException",
		Code:       code,
		Message:    message,
		Data:       struct{}{},
	}
}

// 500 错误处理
func ServerError() *APIException {
	ae := NewAPIException(config.ERROR_SYSTEM, nil)
	ae.HttpStatus = http.StatusInternalServerError
	ae.Tag = "ServerError"
	return ae
}

// 404 错误
func NotFound() *APIException {
	ae := NewAPIException(config.ERROR_NOT_FOUND, nil)
	ae.HttpStatus = http.StatusNotFound
	ae.Message = http.StatusText(http.StatusNotFound)
	ae.Tag = "NotFound"
	return ae
}

// 403 错误
func ForbiddenError() *APIException {
	ae := NewAPIException(config.ERROR_FORBIDDEN, nil)
	ae.HttpStatus = http.StatusForbidden
	ae.Message = http.StatusText(http.StatusForbidden)
	ae.Tag = "ForbiddenError"
	return ae
}

// 参数错误
func ParameterError(message string, err error) *APIException {
	ae := NewAPIException(config.ERROR_PARAMETER, err)
	if len(message) > 1 {
		ae.Message = message
	}
	ae.Tag = "ParameterError"
	return ae
}

// 请求成功
func SuccResp(message string, data interface{}) *APIException {
	ae := NewAPIException(config.SUCC_RESPONSE, nil)
	if len(message) > 1 {
		ae.Message = message
	}
	ae.Data = data
	ae.Tag = "SuccResp"
	return ae
}

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		ae := handler(c)
		if ae.HttpStatus == http.StatusOK {
			c.JSON(ae.HttpStatus, ae)
		}
		if ae.Code > 0 {
			if ae.RawErr != nil {
				Log.Ex(c, ae.Tag, "RawErr: %s, message: %s", ae.RawErr, ae.Message)
			} else {
				Log.Ex(c, ae.Tag, ae.Message)
			}
		}
	}
}
