package scule

// CamelCase formats string by camelCase convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// the normalizeOption parameter for strictly following camelCase
// convention.
//
// Examples:
//
//	CamelCase("foo-bar_baz", nil) // fooBarBaz
//	CamelCase("FooBAR", nil) // fooBAR
//
// The same examples, with normalization:
//
//	CamelCase("foo-bar_baz", &NormalizeOption{ true }) // fooBarBaz
//	CamelCase("FooBAR", &NormalizeOption{ true }) // fooBar
func CamelCase(str string, normalizeOption *NormalizeOption) string {
	if normalizeOption != nil {
		return LowerFirst(PascalCase(str, normalizeOption))
	} else {
		return LowerFirst(PascalCase(str, nil))
	}
}
