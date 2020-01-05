package dependencies_test

import (
	"testing"

	"github.com/psenna/isup/dependencies"
	"github.com/stretchr/testify/assert"
)

func TestGetMockResponse(t *testing.T) {
	var tests = []struct {
		setResponse      dependencies.HTTPResponse
		expectedResponse dependencies.HTTPResponse
		apiMethod        string
		apiURL           string
	}{
		{dependencies.HTTPResponse{}, dependencies.HTTPResponse{RequestMethod: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{RequestMethod: "POST", URL: "localhost:8080/api", StatusCode: 200}, dependencies.HTTPResponse{RequestMethod: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{RequestMethod: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, dependencies.HTTPResponse{RequestMethod: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{RequestMethod: "GET", URL: "localhost:8080/api", StatusCode: 200}, dependencies.HTTPResponse{RequestMethod: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		clientFactory := dependencies.ClientFactory{}

		if test.setResponse.URL != "" {
			clientFactory.AddMockResponse(test.setResponse, test.setResponse.RequestMethod, test.setResponse.URL)
		}

		resultResponse := clientFactory.GetMockResponse(test.apiMethod, test.apiURL)

		assert.True(t, assert.ObjectsAreEqualValues(test.expectedResponse, resultResponse), "The response object are different from the expected response")
	}

}
