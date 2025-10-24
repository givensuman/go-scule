package scule

import (
	"regexp"
	"strings"
)

const titleCaseExceptions string = "(?i)^(a|an|and|as|at|but|by|for|if|in|is|nor|of|on|or|the|to|with)$"

// TitleCase formats a string by capitalizing all words,
// except for minor words. A compact regex of common
// minor words (such as "a", "for", "to") is used.
//
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// the `normalize` parameter for strictly only having the
// first letter uppercased.
//
// Example:
//
// 	TitleCase("this-IS-aTitle") // This is a Title
func TitleCase(str string, normalize bool) string {
	re := regexp.MustCompile(titleCaseExceptions)
	s := SplitByCase(str, nil)

	for i, str := range s {
		if re.Match([]byte(str)) {
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
