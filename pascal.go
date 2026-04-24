package scule

import (
	"strings"
)

// PascalCase formats string by PascalCase convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// the `opts.Normalize = true` parameter for
// strictly following PascalCase convention.
//
// Examples:
//
//	PascalCase("foo-bar_baz", nil) // FooBarBaz
//	PascalCase("FooBAR", nil) 		 // FooBAR
//
// The same examples, with normalization:
//
//	PascalCase("foo-bar_baz", &PascalCaseOptions{ Normalize: true }) // FooBarBaz
//	PascalCase("FooBAR", &PascalCaseOptions{ Normalize: true }) 		 // FooBar
func PascalCase(str string, opts *PascalCaseOptions) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		if opts != nil && opts.Normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, "")
}

type PascalCaseOptions struct {
	Normalize bool
}
