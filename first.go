package scule

import (
	"unicode"
	"unicode/utf8"
)

// UpperFirst converts the first character to
// upper case.
//
// Example:
//
//	UpperFirst("hello world!") // Hello world!
func UpperFirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToUpper(r)) + str[size:]
}

// LowerFirst converts the first character to
// lower case.
//
// Example:
//
//	LowerFirst("Hello world!") // hello world!
func LowerFirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(str)
	return string(unicode.ToLower(r)) + str[size:]
}
