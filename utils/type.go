package utils

import (
	"encoding/json"
	"strconv"
)

// 转换类型为string
func Strval(value interface{}) string {
	var res string
	if value == nil {
		return res
	}

	switch temp := value.(type) {
	case float64:
		res = strconv.FormatFloat(temp, 'f', -1, 64)
	case float32:
		res = strconv.FormatFloat(float64(temp), 'f', -1, 64)
	case int:
		res = strconv.Itoa(temp)
	case uint:
		res = strconv.Itoa(int(temp))
	case int8:
		res = strconv.Itoa(int(temp))
	case uint8:
		res = strconv.Itoa(int(temp))
	case int16:
		res = strconv.Itoa(int(temp))
	case uint16:
		res = strconv.Itoa(int(temp))
	case int32:
		res = strconv.Itoa(int(temp))
	case uint32:
		res = strconv.Itoa(int(temp))
	case int64:
		res = strconv.FormatInt(temp, 10)
	case uint64:
		res = strconv.FormatUint(temp, 10)
	case string:
		res = temp
	case []byte:
		res = string(temp)
	default:
		newValue, _ := json.Marshal(value)
		res = string(newValue)
	}

	return res
}
