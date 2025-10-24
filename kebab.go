package scule

import "strings"

// KebabCase formats string by kebab-case
// convention.
//
// Examples:
//
// 	KebabCase("fooBar_Baz") // foo-bar-baz
// 	KebabCase("foo--bar")		// foo--bar
// 	KebabCase("FooBAR")			// foo-bar
func KebabCase(str string) string {
	return kebabCase(str, "-")
}

func kebabCase(str string, joiner string) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		s[i] = strings.ToLower(str)
	}

	return strings.Join(s, joiner)
}

