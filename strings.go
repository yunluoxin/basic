// Created by East.Zhang at 2022.06.03

package basic

import (
	"strings"
	"unicode/utf8"
)

// IsEmpty 判断字符串是否为空字符串。
func IsEmpty(s string) bool {
	return s == ""
}

// IsBlank 判断是否为空白字符串。
// - 如果为多个空白字符，也算是空白字符串哦！
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) < 1
}

// SizeOfString 获取字符串的真正长度
// Go语言默认的 `len` 是返回字符串的字节长度...
func SizeOfString(s string) int {
	return utf8.RuneCountInString(s)
}