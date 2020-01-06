package dependencies

// HTTPClient HTTPClient
type HTTPClient struct {
	mockEnable   bool
	mockResponse map[string]HTTPResponse
}

// HTTPCall Make a http call
func (c HTTPClient) HTTPCall(request HTTPRequest) HTTPResponse {
	if c.mockEnable {
		return c.GetMockResponse(request.Method, request.URL)
	}

	response := HTTPResponse{}

	switch request.Method {
	case "GET":
		response = c.httpGet(request)
	}

	return response
}

// todo: implementar
func (c HTTPClient) httpGet(request HTTPRequest) HTTPResponse {
	return HTTPResponse{}
}

// AddMockResponse Add a mock response for a api call
func (c *HTTPClient) AddMockResponse(expectedResponse HTTPResponse, apiMethod string, apiURL string) {
	if c.mockResponse == nil {
		c.mockResponse = make(map[string]HTTPResponse)
	}

	c.mockResponse[apiMethod+"-"+apiURL] = expectedResponse
}

// GetMockResponse Get a mock response for a api call
func (c *HTTPClient) GetMockResponse(apiMethod string, apiURL string) HTTPResponse {
	if c.mockResponse == nil {
		return HTTPResponse{RequestMethod: apiMethod, URL: apiURL, StatusCode: 404}
	}

	if _, ok := c.mockResponse[apiMethod+"-"+apiURL]; !ok {
		return HTTPResponse{RequestMethod: apiMethod, URL: apiURL, StatusCode: 404}
	}

	return c.mockResponse[apiMethod+"-"+apiURL]
}

// SetMockEnable Enable or disable mock api call
func (c *HTTPClient) SetMockEnable(enable bool) {
	c.mockEnable = enable
}
