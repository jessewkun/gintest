package config

// 系统错误
const SUCC_RESPONSE = 0
const ERROR_SYSTEM = 10000
const ERROR_NOT_FOUND = 10001
const ERROR_FORBIDDEN = 10002
const ERROR_PARAMETER = 10003
const ERROR_CREATE_REQUEST = 10004
const ERROR_SEND_REQUEST_FAIL = 10005
const ERROR_READ_RESPONSE_FAIL = 10006
const ERROR_NEWCIPHER_FAIL = 10007
const ERROR_DECOCDE_FAIL = 10008

// 业务错误
const ERROR_ADD_USER_FAIL = 20000

var ERROR_MESSAGE_MAP = map[int]string{
	SUCC_RESPONSE:            "succ",
	ERROR_SYSTEM:             "system error",
	ERROR_NOT_FOUND:          "not found",
	ERROR_FORBIDDEN:          "Permission denied",
	ERROR_PARAMETER:          "params is invalid",
	ERROR_CREATE_REQUEST:     "create http request failed",
	ERROR_SEND_REQUEST_FAIL:  "send http request failed",
	ERROR_READ_RESPONSE_FAIL: "read response data failed",
	ERROR_NEWCIPHER_FAIL:     "加密初始化失败",
	ERROR_DECOCDE_FAIL:       "解码失败",

	ERROR_ADD_USER_FAIL: "user add fail",
}
