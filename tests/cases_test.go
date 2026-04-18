package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestUpperCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"foo", "FOO"},
		{"fooBarBaz", "FOO BAR BAZ"},
		{"foo-bar_baz", "FOO BAR BAZ"},
		{"FOO_BAR", "FOO BAR"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.UpperCase(test[0]))
	}
}

func TestLowerCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"FOO", "foo"},
		{"fooBarBaz", "foo bar baz"},
		{"foo-bar_baz", "foo bar baz"},
		{"FOO_BAR", "foo bar"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.LowerCase(test[0]))
	}
}
