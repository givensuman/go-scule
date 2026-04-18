package scule

// DotCase formats a string by dot.case convention.
//
// Example:
//
//	DotCase("fooBarBaz") // foo.bar.baz
func DotCase(str string) string {
	return kebabCase(str, ".")
}
