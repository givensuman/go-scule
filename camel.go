package scule

// CamelCase formats string by camelCase convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// `opts.Normalize = true` for strictly following camelCase
// convention.
//
// Examples:
//
//	CamelCase("foo-bar_baz", nil) // fooBarBaz
//	CamelCase("FooBAR", nil)      // fooBAR
//
// The same examples, with normalization:
//
//	CamelCase("foo-bar_baz", &CamelCaseOptions{ Normalize: true }) // fooBarBaz
//	CamelCase("FooBAR", &CamelCaseOptions{ Normalize: true })      // fooBar
func CamelCase(str string, opts *CamelCaseOptions) string {
	return LowerFirst(PascalCase(str, &PascalCaseOptions{Normalize: opts.Normalize}))
}

type CamelCaseOptions struct {
	Normalize bool
}
