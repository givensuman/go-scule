package scule_test


import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	testCases := [][2]string{
		{"FooBarBaz", "foo_bar_baz"},
		{"FOO_BAR", "foo_bar"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.SnakeCase(test[0]))
	}
}
