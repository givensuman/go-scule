package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestFlatCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "foo"},
		{"foo-bAr", "foobar"},
		{"FooBARb", "foobarb"},
		{"foo_bar-baz/qux", "foobarbazqux"},
		{"FOO_BAR", "foobar"},
		{"foo--bar-Baz", "foobarbaz"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.FlatCase(test[0]))
	}
}
