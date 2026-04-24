package scule

import (
	"strings"
)

// TrainCase formats string by Train-Case
// (a.k.a. HTTP-Header-Case) convention.
// If an uppercase letter is followed by other
// uppercase letters (like WWWAuthenticate),
// they are preserved. You can use the `opts.Normalize = true`
// parameter for strictly only having the first
// letter uppercased.
//
// Example:
//
//	TrainCase("FooBARb") // Foo-Ba-Rb
//	TrainCase("WWWAuthenticate") // WWW-Authenticate
//
// The same examples, with normalization:
//
//	TrainCase("FooBARb", &TrainCaseOptions{ Normalize: true }) // Foo-Barb
//	TrainCase("WWWAuthenticate", &TrainCaseOptions{ Normalize: true }) // Www-Authenticate
func TrainCase(str string, opts *TrainCaseOptions) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		if opts != nil && opts.Normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	s = filterWhitespace(s)
	return strings.Join(s, "-")
}

type TrainCaseOptions struct {
	// Strictly only have the first letter uppercased.
	Normalize bool
}
