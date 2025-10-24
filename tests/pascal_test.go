package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestPascalCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "Foo"},
		{"foo-bAr", "FooBAr"},
		{"FooBARb", "FooBARb"},
		{"foo_bar-baz/qux", "FooBarBazQux"},
		{"FOO_BAR", "FOOBAR"},
		{"foo--bar-Baz", "FooBarBaz"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.PascalCase(test[0], nil))
	}
}

func TestPascalCaseWithNormalization(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "Foo"},
		{"foo-bAr", "FooBAr"},
		{"FooBARb", "FooBaRb"},
		{"foo_bar-baz/qux", "FooBarBazQux"},
		{"FOO_BAR", "FooBar"},
		{"foo--bar-Baz", "FooBarBaz"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.PascalCase(test[0], scule.Options{}.NewNormalize(true)))
	}
}
