package scule

import "strings"

// PascalCase formats a string by PascalCase convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. Pass
// normalize as true to strictly follow PascalCase convention.
//
// Examples:
//
//	PascalCase("foo-bar_baz", false) // FooBarBaz
//	PascalCase("FooBAR", false)      // FooBAR
//
// The same examples, with normalization:
//
//	PascalCase("foo-bar_baz", true) // FooBarBaz
//	PascalCase("FooBAR", true)      // FooBar
func PascalCase(str string, normalize bool) string {
	s := SplitByCase(str)

	for i, str := range s {
		if normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, "")
}
