package scule_test


import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestKebabCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "foo"},
		{"foo/Bar", "foo-bar"},
		{"foo-bAr", "foo-b-ar"},
		{"foo--bar", "foo--bar"},
		{"FooBAR", "foo-bar"},
		{"ALink", "a-link"},
		{"FOO_BAR", "foo-bar"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.KebabCase(test[0], nil))
	}
}
