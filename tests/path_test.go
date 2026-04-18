package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestPathCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"fooBar", "foo/bar"},
		{"foo-bar_baz", "foo/bar/baz"},
		{"FOO_BAR", "foo/bar"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.PathCase(test[0]))
	}
}
