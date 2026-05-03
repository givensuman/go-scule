package scule

import (
	"regexp"
	"strings"
)

var titleCaseExceptionsRe = regexp.MustCompile(`(?i)^(a|an|and|as|at|but|by|for|if|in|is|nor|of|on|or|the|to|with)$`)

// TitleCase formats a string by capitalizing all words except minor
// words (articles, conjunctions, short prepositions). The first word
// is always capitalized regardless.
//
// If uppercase letters are followed by other uppercase letters (like
// FooBAR), they are preserved. Pass normalize as true to strictly
// capitalize only the first letter of each word.
//
// Examples:
//
//	TitleCase("this-IS-aTitle", false) // This is a Title
//	TitleCase("THIS is a TITLE", false) // THIS is a TITLE
//
// The same examples, with normalization:
//
//	TitleCase("this-IS-aTitle", true) // This is a Title
//	TitleCase("THIS is a TITLE", true) // This is a Title
func TitleCase(str string, normalize bool) string {
	s := SplitByCase(str)

	for i, str := range s {
		if i > 0 && titleCaseExceptionsRe.MatchString(str) {
			s[i] = strings.ToLower(str)
			continue
		}

		if normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, " ")
}
