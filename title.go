package scule

import (
	"regexp"
	"strings"
)

const titleCaseExceptions string = "(?i)^(a|an|and|as|at|but|by|for|if|in|is|nor|of|on|or|the|to|with)$"

var titleCaseExceptionsRe = regexp.MustCompile(titleCaseExceptions)

// TitleCase formats a string by capitalizing all words,
// except for minor words. A compact regex of common
// minor words (such as "a", "for", "to") is used.
// The first word is always capitalized regardless.
//
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// `opts.Normalize = true` for strictly only having the
// first letter uppercased.
//
// Example:
//
//	TitleCase("this-IS-aTitle") // This is a Title
func TitleCase(str string, opts *TitleCaseOptions) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		if i > 0 && titleCaseExceptionsRe.MatchString(str) {
			s[i] = strings.ToLower(str)
			continue
		}

		if opts != nil && opts.Normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, " ")
}

type TitleCaseOptions struct {
	// Strictly only have the first letter uppercased.
	Normalize bool
}
