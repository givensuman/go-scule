package scule

import "strings"

// TrainCase formats a string by Train-Case (HTTP-Header-Case) convention.
// If uppercase letters are followed by other uppercase letters (like
// WWWAuthenticate), they are preserved. Pass normalize as true to
// strictly capitalize only the first letter of each word.
//
// Examples:
//
//	TrainCase("FooBARb", false)        // Foo-BA-Rb
//	TrainCase("WWWAuthenticate", false) // WWW-Authenticate
//
// The same examples, with normalization:
//
//	TrainCase("FooBARb", true)        // Foo-Barb
//	TrainCase("WWWAuthenticate", true) // Www-Authenticate
func TrainCase(str string, normalize bool) string {
	s := SplitByCase(str)

	for i, str := range s {
		if normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	s = filterWhitespace(s)
	return strings.Join(s, "-")
}
