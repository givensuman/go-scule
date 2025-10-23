package scule

import (
	"strings"
)

/*
PascalCase formats string by PascalCase convention.

If an uppercase letter is followed by other uppercase
letters (like FooBAR), they are preserved. You can use
the `normalize` parameter for strictly following PascalCase
convention.

Example:

	PascalCase("foo-bar_baz", nil) 		// FooBarBaz
	PascalCase("FooBAR", nil) 		 		// FooBAR
	PascalCase("FooBAR", nil) 			 	// FooBar
*/
func PascalCase(str string, opts *Options) string {
	if opts == nil {
		opts = defaultOptions()
	}

	s := SplitByCase(str, nil)

	for i, str := range s {
		if opts.Normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, "")
}
