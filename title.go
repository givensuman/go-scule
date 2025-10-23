package scule

import (
	"regexp"
	"strings"
)

const titleCaseExceptions string = "(?i)^(a|an|and|as|at|but|by|for|if|in|is|nor|of|on|or|the|to|with)$"

/*
TitleCase formats a string by capitalizing all words,
except for minor words. A compact regex of common
minor words (such as "a", "for", "to") is used.

Example:

	TitleCase("this-IS-aTitle") // This is a Title
*/
func TitleCase(str string, normalize *Normalize) string {
	if normalize == nil {
		normalize = Options{}.NewNormalize(true)
	}

	re := regexp.MustCompile(titleCaseExceptions)
	s := SplitByCase(str, nil)

	for i, str := range s {
		if re.Match([]byte(str)) {
			s[i] = strings.ToLower(str)
			continue
		}

		if normalize.isTrue() {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, " ")
}
