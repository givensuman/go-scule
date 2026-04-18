package scule

// PathCase formats a string by path/case convention.
//
// Example:
//
//	PathCase("fooBarBaz") // foo/bar/baz
func PathCase(str string) string {
	return kebabCase(str, "/")
}
