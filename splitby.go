package scule

import (
	"slices"
	"strings"
	"unicode"
)

// SplitByCase splits a string by the splitters provided
// (default: {"-", "_", "/", ".", " "}). It also splits by
// case change of upper-to-lower or lower-to-upper. Numbers
// are ignored for case changes.
//
// Case is preserved in the returned slice, and splitters
// are omitted.
//
// Examples:
//
//	SplitByCase("foo-bar_baz", nil) // { "foo", "bar", "baz" }
//	SplitByCase("fooBarBaz", nil)   // { "foo", "Bar", "Baz" }
//	SplitByCase("FOOBar", nil)      // { "FOO", "Bar" }
//	SplitByCase("foo123-bar", nil)  // { "foo123", "bar" }
//
// Example with custom splitters:
//
//	SplitByCase("foo//Bar.fizBaz", &Splitters{ "//", "." })
//	// { "foo", "Bar", "fiz", "Baz" }
func SplitByCase(str string, splitters *Splitters) []string {
	if splitters == nil {
		splitters = &Splitters{"-", "_", "/", ".", " "}
	}

	if len(str) == 0 {
		return []string{}
	}

	var parts []string
	builder := strings.Builder{}

	var previousWasUpper bool
	var previousWasSplitter bool

	runes := []rune(str)
	for i, c := range runes {
		isUpper := unicode.IsUpper(c)
		isSplitter := slices.Contains(splitters.toSlice(), string(c))

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
				builder.WriteRune(c)

				previousWasUpper = isUpper
				previousWasSplitter = false
				continue
			}

			// Falling edge case
			if previousWasUpper && !isUpper && builder.Len() > 1 {
				s := []rune(builder.String())
				lastRune := s[len(s)-1]

				parts = append(parts, string(s[:len(s)-1]))

				builder.Reset()
				builder.WriteRune(lastRune)
				builder.WriteRune(c)

				previousWasUpper = isUpper
				previousWasSplitter = false
				continue
			}
		}

		// Normal character
		builder.WriteRune(c)
		previousWasUpper = isUpper
		previousWasSplitter = false
	}

	parts = append(parts, builder.String())
	return parts
}

type Splitters []string

func (s *Splitters) toSlice() []string {
	return []string(*s)
}
