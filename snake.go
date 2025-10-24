package scule

// SnakeCase formats string by snake_case
// convention.
//
// Example:
//
//	SnakeCase("foo-barBaz") // foo_bar_baz
func SnakeCase(str string) string {
	return kebabCase(str, "_")
}
