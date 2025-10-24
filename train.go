package scule

import (
	"strings"
)

// TrainCase formats string by Train-Case
// (a.k.a. HTTP-Header-Case) convention.
//
// if an uppercase letter is followed by other
// uppercase letters (like WWWAuthenticate),
// they are preserved. You can use the `normalize`
// parameter for strictly only having the first
// letter uppercased.
//
// Example:
//
// 	TrainCase("FooBARb") // Foo-Ba-Rb
// 	TrainCase("WWWAuthenticate") // WWW-Authenticate
func TrainCase(str string, normalize bool) string {
	s := SplitByCase(str, nil)

	for i, str := range s {
		if normalize {
			str = strings.ToLower(str)
		}

		s[i] = UpperFirst(str)
	}

	s = filterWhitespace(s)
	return strings.Join(s, "-")
}

func filterWhitespace(strs []string) []string {
	var out []string
	
	for _, str := range strs {
		if str == "" || str == " " {
			continue
		}

		out = append(out, str)
	}

	return out
}
