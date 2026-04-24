package scule

import (
	"strings"
)

// PascalCase formats string by PascalCase convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// the normalizeOption parameter for strictly following
// PascalCase convention.
//
// Examples:
//
//	PascalCase("foo-bar_baz", nil) // FooBarBaz
//	PascalCase("FooBAR", nil) // FooBAR
//
// The same examples, with normalization:
//
//	PascalCase("foo-bar_baz", &NormalizeOption{ true }) // FooBarBaz
//	PascalCase("FooBAR", &NormalizeOption{ true }) // FooBar
func PascalCase(str string, normalizeOption *NormalizeOption) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		if normalizeOption != nil && normalizeOption.Normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, "")
}
