package lib

import (
	"unicode"
	"unicode/utf8"
)

func UpperFirst(s string) (_ string) {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
