package scule

import "strings"

// UpperCase formats a string by splitting into words and
// joining them with spaces, converting all characters to
// upper case.
//
// Example:
//
//	UpperCase("fooBarBaz") // FOO BAR BAZ
func UpperCase(str string) string {
	s := filterWhitespace(SplitByCase(str, nil))
	for i, w := range s {
		s[i] = strings.ToUpper(w)
	}
	return strings.Join(s, " ")
}

// LowerCase formats a string by splitting into words and
// joining them with spaces, converting all characters to
// lower case.
//
// Example:
//
//	LowerCase("fooBarBaz") // foo bar baz
func LowerCase(str string) string {
	s := filterWhitespace(SplitByCase(str, nil))
	for i, w := range s {
		s[i] = strings.ToLower(w)
	}
	return strings.Join(s, " ")
}
