package scule

import (
	"slices"
	"strings"
	"unicode"
)

/*
SplitByCase splits a string by the splitters provided
(default: {'-', '_', '/', '.'}). It also splits by
case change of upper-to-lower or lower-to-upper. Numbers
are ignored for case changes.

Case is preserved in the returned slice, and splitters
are omitted.

Example:

	SplitByCase("foo-bar_baz") // ["foo", "bar", "baz"]
	SplitByCase("fooBarBaz")   // ["foo", "Bar", "Baz"]
	SplitByCase("FOOBar")			 // ["FOO", "Bar"]
	SplitByCase("foo123-bar")	 // ["foo123", "bar"]

Example with custom splitters:

	SplitByCase("foo//Bar.fizBaz", &[]byte{"//"})
	// ["foo", "Bar.fiz", "Baz"]
*/
func SplitByCase(str string, splitters *[]byte) []string {
	if splitters == nil {
		splitters = &[]byte{'-', '_', '/', '.'}
	}

	if len(str) == 0 {
		return []string{}
	}

	var parts []string
	builder := strings.Builder{}

	var previousWasUpper bool
	var previousWasSplitter bool

	for i := range str {
		c := str[i]
		isUpper := unicode.IsUpper(rune(str[i]))
		isSplitter := slices.Contains(*splitters, c)

		// Splitter case
		if isSplitter {
			parts = append(parts, builder.String())
			builder.Reset()

			previousWasUpper = false
			previousWasSplitter = true
			continue
		}

		if !previousWasSplitter {
			// Rising edge case
			if !previousWasUpper && isUpper && i > 0 {
				parts = append(parts, builder.String())

				builder.Reset()
				builder.WriteByte(c)

				previousWasUpper = isUpper
				previousWasSplitter = false
				continue
			}

			// Falling edge case
			if previousWasUpper && !isUpper && builder.Len() > 1 {
				s := builder.String()
				lastChar := s[len(s)-1]

				parts = append(parts, s[:len(s)-1])

				builder.Reset()
				builder.WriteByte(lastChar)
				builder.WriteByte(c)

				previousWasUpper = isUpper
				previousWasSplitter = false
				continue
			}
		}

		// Normal character
		builder.WriteByte(c)
		previousWasUpper = isUpper
		previousWasSplitter = false
	}

	parts = append(parts, builder.String())
	return parts
}
