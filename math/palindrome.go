package math

import (
	"strconv"
)

func IsPalindrome(n int) bool {
	str := strconv.Itoa(n)
	var revStr string
	for _, c := range str {
		revStr = string(c) + revStr
	}
	return str == revStr
}
