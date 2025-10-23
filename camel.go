package scule

/*
CamelCase formats string by camelCase
convention.

Example:
	
	CamelCase("foo-bar_baz") // fooBarBaz
*/
func CamelCase(str string, opts *Options) string {
	if opts == nil {
		opts = defaultOptions()
	}

	return LowerFirst(PascalCase(str, opts))
}

