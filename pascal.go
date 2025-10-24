package scule

import (
	"strings"
)

// PascalCase formats string by PascalCase convention.
//
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// the `normalize` parameter for
// strictly following PascalCase convention.
//
// Examples:
//
// 	PascalCase("foo-bar_baz", nil) // FooBarBaz
// 	PascalCase("FooBAR", nil) 		 // FooBAR
//
// The same examples, with normalization:
//
// 	PascalCase("foo-bar_baz", Options{}.NewNormalize(true)) // FooBarBaz
// 	PascalCase("FooBAR", Options{}.NewNormalize(true)) 			// FooBar
func PascalCase(str string, normalize bool) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		if normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, "")
}
