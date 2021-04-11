package util

import (
	"regexp"
)

// 验证邮箱格式
func CheckEmail(email string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.(com|cn|org|dev)$", email)
	return match
}

// 截取字符串 并返回长度
func StrSplice(str string, start, end int) string {
	if len(str) == 0 {
		return str
	}

	s := []rune(str)

	if start == -1 {
		return string(s[:end])
	}

	if end == -1 {
		return string(s[start:])
	}

	return string(s[start:end])
}
