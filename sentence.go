package scule

import "strings"

// SentenceCase formats a string by capitalizing only the
// first word, joining all words with spaces in lower case.
//
// Example:
//
//	SentenceCase("fooBarBaz") // Foo bar baz
func SentenceCase(str string) string {
	s := filterWhitespace(SplitByCase(str, nil))
	for i, w := range s {
		if i == 0 {
			s[i] = UpperFirst(strings.ToLower(w))
		} else {
			s[i] = strings.ToLower(w)
		}
	}
	return strings.Join(s, " ")
}
