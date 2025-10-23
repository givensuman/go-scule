package scule

import "strings"

/*
KebabCase formats string by kebab-case
convention.

Example:

	KebabCase("fooBar_Baz") // foo-bar-baz
*/
func KebabCase(str string, joiner *Joiner) string {
	if joiner == nil {
		joiner = Options{}.NewJoiner("-")
	}

	s := SplitByCase(str, nil)

	for i, str := range s {
		s[i] = strings.ToLower(str)
	}

	return strings.Join(s, joiner.Value)
}

