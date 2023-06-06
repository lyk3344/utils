package utils

import (
	"errors"
	"reflect"
)

var (
	ErrUnSupportSlice = errors.New("haystack must be slice, array")
)

func FindInSlice(haystack []string, val string) bool {
	for _, v := range haystack {
		if v == val {
			return true
		}
	}

	return false
}

func FindInSliceByInterface(haystack interface{}, val interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == val {
				return true, nil
			}
		}

		return false, nil
	}

	return false, ErrUnSupportSlice
}

//通过循环来去重
//slice元素为string
func RemoveDuplicateStrByLoop(s []string) []string {
	var (
		result []string
		flag bool
	)
	for _, i := range s {
		flag = true
		for _, j := range result {
			if i == j {
				flag = false
				break
			}
		}
		if flag {
			s = append(result, i)
		}
	}
	return result
}

//通过循环来去重
//slice元素为int
func RemoveDuplicateIntByLoop(s []int) []int {
	var (
		result []int
		flag bool
	)
	for _, i := range s {
		flag = true
		for _, j := range result {
			if i == j {
				flag = false
				break
			}
		}
		if flag {
			s = append(result, i)
		}
	}
	return result
}

//通过循环来去重
//slice元素为int64
func RemoveDuplicateInt64ByLoop(s []int64) []int64 {
	var (
		result []int64
		flag bool
	)
	for _, i := range s {
		flag = true
		for _, j := range result {
			if i == j {
				flag = false
				break
			}
		}
		if flag {
			s = append(result, i)
		}
	}
	return result
}

//通过map来对slice进行去重
//slice元素为string
func RemoveDuplicateStrByMap(s []string) []string {
	var (
		result []string
		uniqMap map[string]int
	)
	for _, i := range s {
		l := len(uniqMap)
		uniqMap[i] = 0
		if len(uniqMap) != l {	//如果有新的key出来，说明map变长了
			result = append(result, i)
		}
	}
	return result
}

//通过map来对slice进行去重
//slice元素为int
func RemoveDuplicateIntByMap(s []int) []int {
	var (
		result []int
		uniqMap map[int]int
	)
	for _, i := range s {
		l := len(uniqMap)
		uniqMap[i] = 0
		if len(uniqMap) != l {	//如果有新的key出来，说明map变长了
			result = append(result, i)
		}
	}
	return result
}

//通过map来对slice进行去重
//slice元素为int64
func RemoveDuplicateInt64ByMap(s []int64) []int64 {
	var (
		result []int64
		uniqMap map[int64]int
	)
	for _, i := range s {
		l := len(uniqMap)
		uniqMap[i] = 0
		if len(uniqMap) != l {	//如果有新的key出来，说明map变长了
			result = append(result, i)
		}
	}
	return result
}

//高效去重
//slice元素为string类型
func RemoveDuplicateStrCellInSlice(s []string) []string {
	if len(s) < 1024 {
		return RemoveDuplicateStrByLoop(s)
	} else {
		return RemoveDuplicateStrByMap(s)
	}
}

//高效去重
//slice元素为int类型
func RemoveDuplicateIntCellInSlice(s []int) []int {
	if len(s) < 1024 {
		return RemoveDuplicateIntByLoop(s)
	} else {
		return RemoveDuplicateIntByMap(s)
	}
}

//高效去重
//slice元素为int64类型
func RemoveDuplicateInt64CellInSlice(s []int64) []int64 {
	if len(s) < 1024 {
		return RemoveDuplicateInt64ByLoop(s)
	} else {
		return RemoveDuplicateInt64ByMap(s)
	}
}

