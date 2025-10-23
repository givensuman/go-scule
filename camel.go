package scule

/*
CamelCase formats string by camelCase
convention.

Example:
	
	CamelCase("foo-bar_baz") // fooBarBaz
*/
func CamelCase(str string, normalize *Normalize) string {
	if normalize == nil {
		normalize = Options{}.NewNormalize(true)
	}

	return LowerFirst(PascalCase(str, normalize))
}

