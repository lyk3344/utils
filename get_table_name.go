package utils

import (
	"strconv"
)

const (
	segmentAndOpVal = 256
)

// 通过表前缀及表分片信息获取对应的表名
func GetName(table, index string) (bool, string) {
	defer func() {
		if e := recover(); e != nil { //发生宕机时，获取panic传递的上下文并打印
			//fmt.Println("index太短")
			panic("index太短")
		}
	}()
	if len(index) < 4 {
		return true, table
	}
	ix := int(CityHash32([]byte(index), uint32(16)) % segmentAndOpVal)
	index = strconv.Itoa(ix)
	if ix < 10 {
		index = "00" + index
	} else if ix < 100 {
		index = "0" + index
	}
	return ix < 128, table + "_" + index
}