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
		{"éclair", "Éclair"},
		{"über", "Über"},
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
		{"Éclair", "éclair"},
		{"Über", "über"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.LowerFirst(test[0]))
	}
}
