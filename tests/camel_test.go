package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestCamelCase(t *testing.T) {
	testCases := [][2]string{
		{"FooBarBaz", "fooBarBaz"},
		{"FOO_BAR", "fooBar"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.CamelCase(test[0], true))
	}
}
