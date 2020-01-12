package dependencies_test

import (
	"github.com/psenna/isup/dependencies"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	timeOut    = 2000
	minTimeOut = 50
	maxTimeout = 60000
)

func TestGetHTTPRequestTimeOut(t *testing.T) {
	var tests = []struct {
		apiMethod       string
		apiURL          string
		timeOut         int
		expectedTimeout int
	}{
		{"GET", "localhost:8080/api", 1, minTimeOut},
		{"GET", "localhost:8080/api", 50, 50},
		{"GET", "localhost:8080/api", 0, timeOut},
		{"GET", "localhost:8080/api", 60000, 60000},
		{"GET", "localhost:8080/api", 60001, maxTimeout},
	}

	for _, test := range tests {
		request := dependencies.GetHTTPRequest(test.apiMethod, test.apiURL)

		if test.timeOut != 0 {
			request = request.SetTimeOut(test.timeOut)
		}

		assert.Equal(t, test.expectedTimeout, request.GetTimeOut())
	}
}
