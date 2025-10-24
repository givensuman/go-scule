package scule_test


import (
	"testing"

	"github.com/givensuman/go-scule"
	"github.com/stretchr/testify/assert"
)

func TestTitleCase(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"f", "F"},
		{"foo", "Foo"},
		{"foo-bar", "Foo Bar"},
		{"this-IS-aTitle", "This is a Title"},
		{"THIS is a TITLE", "THIS is a TITLE"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.TitleCase(test[0], false))
	}
}

func TestTitleCaseWithNormalize(t *testing.T) {
	testCases := [][2]string{
		{"", ""},
		{"f", "F"},
		{"foo", "Foo"},
		{"foo-bar", "Foo Bar"},
		{"this-IS-aTitle", "This is a Title"},
		{"THIS is a TITLE", "This is a Title"},
	}

	for _, test := range testCases {
		assert.Equal(t, test[1], scule.TitleCase(test[0], true))
	}
}
