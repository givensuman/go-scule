package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestSplitByCase(t *testing.T) {
	testCases := [][][]string{
		{{""}, nil},
		{{"foo"}, {"foo"}},
		{{"fooBar"}, {"foo", "Bar"}},
		{{"FooBarBaz"}, {"Foo", "Bar", "Baz"}},
		{{"FooBARb"}, {"Foo", "BA", "Rb"}},
		{{"foo_bar-baz/qux"}, {"foo", "bar", "baz", "qux"}},
		{{"foo--bar-Baz"}, {"foo", "", "bar", "Baz"}},
		{{"FOO_BAR"}, {"FOO", "BAR"}},
		{{"foo123-bar"}, {"foo123", "bar"}},
		{{"FOOBar"}, {"FOO", "Bar"}},
		{{"ALink"}, {"A", "Link"}},
		{{"héllo-wörld"}, {"héllo", "wörld"}},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.SplitByCase(test[0][0]))
	}
}

func TestSplitByCaseWithCustomSplitter(t *testing.T) {
	assert.Equal(t,
		[]string{"foo", "Bar", "fuzz", "FI", "Zz"},
		scule.SplitByCase("foo\\Bar.fuzz-FIZz", "\\", ".", "-"),
	)
	assert.Equal(t,
		[]string{"new-name-value"},
		scule.SplitByCase("new-name-value", "_"),
	)
}
