package scule_test

import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestSentenceCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"fooBar", "Foo bar"},
		{"foo-bar_baz", "Foo bar baz"},
		{"FOO_BAR", "Foo bar"},
		{"this-IS-aTitle", "This is a title"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.SentenceCase(test[0]))
	}
}
