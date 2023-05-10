package utils

import (
	"fmt"
	"strings"
)

func Split(r rune) bool {
	return r == '<' || r == '>'
	//return r == '+' || r == '-'||r == '*'||r == '/'||r == '+'||r == '('||r == ')'||r == '%'||r=='['||r==']'||r=='^'
}
func SplitString(s string) []string {
	a := strings.FieldsFunc(s, Split)
	return a
}

//将slice转换成以逗号分割的字符串
func Convert(array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
