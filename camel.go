package scule

// CamelCase formats a string by camelCase convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. Pass
// normalize as true to strictly follow camelCase convention.
//
// Examples:
//
//	CamelCase("foo-bar_baz", false) // fooBarBaz
//	CamelCase("FooBAR", false)      // fooBAR
//
// The same examples, with normalization:
//
//	CamelCase("foo-bar_baz", true) // fooBarBaz
//	CamelCase("FooBAR", true)      // fooBar
func CamelCase(str string, normalize bool) string {
	return LowerFirst(PascalCase(str, normalize))
}
