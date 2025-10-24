package scule

import "strings"

// UpperFirst converts the first character to
// upper case.
//
// Example:
//
// 	UpperFirst("hello world!") // Hello world!
func UpperFirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	return strings.ToUpper(string(str[0])) + str[1:]
}

// LowerFirst converts the first character to
// lower case.
//
// Example:
//
// 	LowerFirst("Hello world!") // hello world!
func LowerFirst(str string) string {
	if len(str) == 0 {
		return ""
	}

	return strings.ToLower(string(str[:1])) + str[1:]
}
