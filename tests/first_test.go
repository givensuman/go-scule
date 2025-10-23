package scule_test


import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestUpperFirst(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "Foo"},
		{"Foo", "Foo"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.UpperFirst(test[0]))
	}
}

func TestLowerFirst(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "foo"},
		{"Foo", "foo"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.LowerFirst(test[0]))
	}
}
