package scule

// KebabCase formats string by kebab-case
// convention.
//
// Examples:
//
//	KebabCase("fooBar_Baz") // foo-bar-baz
//	KebabCase("foo--bar") // foo--bar
//	KebabCase("FooBAR") // foo-bar
func KebabCase(str string) string {
	return joinWith(str, "-")
}
