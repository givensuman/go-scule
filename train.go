package scule

import (
	"strings"
)

/*
TrainCase formats string by Train-Case
(a.k.a. HTTP-Header-Case) convention.

if an uppercase letter is followed by other
uppercase letters (like WWWAuthenticate),
they are preserved. You can use the `normalize`
parameter for strictly only having the first
letter uppercased.

Example:

	TrainCase("FooBARb") // Foo-Ba-Rb
	TrainCase("WWWAuthenticate") // WWW-Authenticate
*/
func TrainCase(str string, normalize *Normalize) string {
	if normalize == nil {
		normalize = Options{}.NewNormalize(true)
	}

	s := SplitByCase(str, nil)

	for i, str := range s {
		if normalize.isTrue() {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	return strings.Join(s, "-")
}
