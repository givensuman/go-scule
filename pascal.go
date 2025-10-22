package scule

/*
PascalCase formats string by PascalCase convention.

If an uppercase letter is followed by other uppercase
letters (like FooBAR), they are preserved. You can use
the `normalize` parameter for strictly following PascalCase
convention.

Example:

	PascalCase("foo-bar_baz", false) // FooBarBaz
	PascalCase("FooBAR", false) 		 // FooBAR
	PascalCase("FooBAR", true) 			 // FooBar
*/
func PascalCase(str string, normalize bool) string {
	// TODO
	return ""
}
