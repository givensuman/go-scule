package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestSplitByCase(t *testing.T) {
	testCases := [][][]string{
		{{""}, {}},
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
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.SplitByCase(test[0][0], nil))
	}
}

func TestSplitByCaseWithCustomSplitter(t *testing.T) {
	testCases := [][][]string{
		{{"foo\\Bar.fuzz-FIZz"}, {"foo", "Bar", "fuzz", "FI", "Zz"}},
		{{"new-name-value"}, {"new-name-value"}},
	}

	testSplitters := [][]byte{
		{'\\', '.', '-'},
		{'_'},
	}

	for i, test := range testCases {
		assert.Equal(t, test[1], scule.SplitByCase(test[0][0], &testSplitters[i]))
	}
}
