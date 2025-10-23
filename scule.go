/*
Package scule provides string case utiliites.
*/
package scule

/*
Options provides additional options to
string case conversion.
*/
type Options struct{}

// Normalize dictates whether
// to strictly follow case conventions.
type Normalize struct {
	Value bool
}

func (n *Normalize) isTrue() bool {
	return n.Value
}

// NewNormalize creates a new Normalize instance.
func (o Options) NewNormalize(b bool) *Normalize {
	return &Normalize{Value: b}
}

// Joiner defines byte by which to join
// formatting string.
type Joiner struct {
	Value string
}

// NewJoiner creates a new Joiner instance.
func (o Options) NewJoiner(b string) *Joiner {
	return &Joiner{Value: b}
}

func defaultSplitters() *[]string {
	return &[]string{"-", "_", "/", "."}
}
