package scule

// FlatCase formats string by flatcase
// convention.
//
// Example:
//
// 	FlatCase("foo-barBaz") // foobarbaz
func FlatCase(str string) string {
	return kebabCase(str, "")
}
