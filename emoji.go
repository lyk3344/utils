package utils

import "strconv"

//将字符串中的emoji转换为unicode编码
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `\u` + strconv.FormatInt(int64(rs[i]), 16) + ``
			ret += u
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}