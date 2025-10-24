package scule

// CamelCase formats string by camelCase
// convention.
// If an uppercase letter is followed by other uppercase
// letters (like FooBAR), they are preserved. You can use
// the `normalize` parameter (defaults to false) for 
// strictly following PascalCase convention.
// 
// Examples:
//
// 	PascalCase("foo-bar_baz", nil) // fooBarBaz
// 	PascalCase("FooBAR", nil) 		 // fooBAR
//
// The same examples, with normalization:
//
// 	PascalCase("foo-bar_baz", Options{}.NewNormalize(true)) // fooBarBaz
// 	PascalCase("FooBAR", Options{}.NewNormalize(true)) 			// fooBar
func CamelCase(str string, normalize bool) string {
	return LowerFirst(PascalCase(str, normalize))
}

