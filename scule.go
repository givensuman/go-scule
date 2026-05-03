// Package scule provides string case utilities.
package scule

import "strings"

func filterWhitespace(strs []string) []string {
	var out []string

	for _, str := range strs {
		if strings.TrimSpace(str) == "" {
			continue
		}

		out = append(out, str)
	}

	return out
}

func joinWith(str string, joiner string) string {
	s := filterWhitespace(SplitByCase(str))

	for i, str := range s {
		s[i] = strings.ToLower(str)
	}

	return strings.Join(s, joiner)
}
