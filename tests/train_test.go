package scule_test

// import (
// 	"testing"
//
// 	"github.com/givensuman/go-scule"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestTrainCase(t *testing.T) {
// 	testCases := [][2]string{
// 		{"", ""},
// 		{"f", "F"},
// 		{"foo", "Foo"},
// 		{"foo-bAr", "Foo-B-Ar"},
// 		{"AcceptCH", "Accept-CH"},
// 		{"foo_bar-baz/qux", "Foo-Bar-Baz-Qux"},
// 		{"FOO_BAR", "FOO-BAR"},
// 		{"foo--bar-Baz", "Foo-Bar-Baz"},
// 		{"WWW-authenticate", "WWW-Authenticate"},
// 		{"WWWAuthenticate", "WWW-Authenticate"},
// 	}
//
// 	for _, test := range testCases {
// 		assert.Equal(t, test[1], scule.TrainCase(test[0], &scule.Options{Normalize: false}))
// 	}
// }
//
// func TestTrainCaseWithNormalize(t *testing.T) {
// 	testCases := [][2]string{
// 		{"AcceptCH", "Accept-Ch"},
// 		{"FOO_BAR", "Foo-Bar"},
// 		{"WWW-authenticate", "Www-Authenticate"},
// 	}
//
// 	for _, test := range testCases {
// 		assert.Equal(t, test[1], scule.TrainCase(test[0], nil))
// 	}
// }
