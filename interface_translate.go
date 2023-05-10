package utils

import (
	"encoding/json"
	"strconv"
)

func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}

func GetInterfaceToInt64(t1 interface{}) int64 {
	var t2 int64
	switch t1.(type) {
	case uint:
		t2 = int64(t1.(uint))
		break
	case int8:
		t2 = int64(t1.(int8))
		break
	case uint8:
		t2 = int64(t1.(uint8))
		break
	case int16:
		t2 = int64(t1.(int16))
		break
	case uint16:
		t2 = int64(t1.(uint16))
		break
	case int32:
		t2 = int64(t1.(int32))
		break
	case uint32:
		t2 = int64(t1.(uint32))
		break
	case uint64:
		t2 = int64(t1.(uint64))
		break
	case float32:
		t2 = int64(t1.(float32))
		break
	case float64:
		t2 = int64(t1.(float64))
		break
	case string:
		t2, _ = strconv.ParseInt(t1.(string), 10, 64)
		break
	default:
		t2 = t1.(int64)
		break
	}
	return t2
}

// interface 转 string
func GetInterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}