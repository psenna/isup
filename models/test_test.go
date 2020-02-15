package models_test

import (
	"testing"

	"github.com/psenna/isup/models"
	"github.com/stretchr/testify/assert"
)

func TestTestIsValid(t *testing.T) {
	var tests = []struct {
		t        models.Test
		valid    bool
		problems string
	}{
		{models.Test{}, false, "A test must have a name\nThe type  is invalid\nA test must have a type\nA test must have a url\nA test must have a system\n"},
		{models.Test{Name: "a test"}, false, "The type  is invalid\nA test must have a type\nA test must have a url\nA test must have a system\n"},
		{models.Test{Name: "a test", TestType: "Any"}, false, "The type Any is invalid\nA test must have a url\nA test must have a system\n"},
		{models.Test{Name: "a test", TestType: "HTTP-GET"}, false, "A test must have a url\nA test must have a system\n"},
		{models.Test{Name: "a test", TestType: "Any", System: "Github", URL: "http://github.com"}, false, "The type Any is invalid\n"},
		{models.Test{Name: "a test", TestType: "HTTP-GET", System: "Github", URL: "http://github.com"}, true, ""},
		{models.Test{Name: "a test", TestType: "HTTP-POST", System: "Github", URL: "http://github.com"}, true, ""},
	}

	for _, test := range tests {
		valid, problems := test.t.Validate()

		assert.Equal(t, valid, test.valid)
		assert.Equal(t, test.t.Valid, test.valid)
		assert.Equal(t, problems, test.problems)
	}
}
