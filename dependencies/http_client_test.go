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
		{dependencies.HTTPResponse{}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{Method: "POST", URL: "localhost:8080/api", StatusCode: 200}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		HTTPClient := dependencies.HTTPClient{}

		if test.setResponse.URL != "" {
			HTTPClient.AddMockResponse(test.setResponse, test.setResponse.Method, test.setResponse.URL)
		}

		resultResponse := HTTPClient.GetMockResponse(test.apiMethod, test.apiURL)

		assert.True(t, assert.ObjectsAreEqualValues(test.expectedResponse, resultResponse), "The response object are different from the expected response")
	}
}

// Get a response with mock enable
func TestGetResponseMockEnable(t *testing.T) {
	var tests = []struct {
		setResponse      dependencies.HTTPResponse
		expectedResponse dependencies.HTTPResponse
		apiMethod        string
		apiURL           string
	}{
		{dependencies.HTTPResponse{}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{Method: "POST", URL: "localhost:8080/api", StatusCode: 200}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/v2/api", StatusCode: 200}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 404}, "GET", "localhost:8080/api"},
		{dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, dependencies.HTTPResponse{Method: "GET", URL: "localhost:8080/api", StatusCode: 200}, "GET", "localhost:8080/api"},
	}

	for _, test := range tests {
		HTTPClient := dependencies.HTTPClient{}
		HTTPClient.SetMockEnable(true)

		if test.setResponse.URL != "" {
			HTTPClient.AddMockResponse(test.setResponse, test.setResponse.Method, test.setResponse.URL)
		}

		resultResponse := HTTPClient.HTTPCall(dependencies.GetHTTPRequest(test.apiMethod, test.apiURL))

		assert.True(t, assert.ObjectsAreEqualValues(test.expectedResponse, resultResponse), "The response object are different from the expected response")
	}
}

// Get response from a api
func TestGetResponseFromApi(t *testing.T) {
	HTTPClient := dependencies.HTTPClient{}

	httpRequest := dependencies.GetHTTPRequest("GET", "https://github.com/")

	response := HTTPClient.HTTPCall(httpRequest)

	assert.Equal(t, 200, response.StatusCode)
}
