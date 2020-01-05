package dependencies

// HTTPResponse A response from a http call
type HTTPResponse struct {
	URL           string
	RequestMethod string
	StatusCode    int
	Body          string
	ResponseTime  float64
	ContentLength int64
	ContentType   string
}

// ClientFactory ClientFactory
type ClientFactory struct {
	mockEnable   bool
	mockResponse map[string]HTTPResponse
}

// AddMockResponse Add a mock response for a api call
func (c *ClientFactory) AddMockResponse(expectedResponse HTTPResponse, apiMethod string, apiURL string) {
	if c.mockResponse == nil {
		c.mockResponse = make(map[string]HTTPResponse)
	}

	c.mockResponse[apiMethod+"-"+apiURL] = expectedResponse
}

// GetMockResponse Get a mock response for a api call
func (c *ClientFactory) GetMockResponse(apiMethod string, apiURL string) HTTPResponse {
	if c.mockResponse == nil {
		return HTTPResponse{RequestMethod: apiMethod, URL: apiURL, StatusCode: 404}
	}

	if _, ok := c.mockResponse[apiMethod+"-"+apiURL]; !ok {
		return HTTPResponse{RequestMethod: apiMethod, URL: apiURL, StatusCode: 404}
	}

	return c.mockResponse[apiMethod+"-"+apiURL]
}

// SetMockEnable Enable or disable mock api call
func (c *ClientFactory) SetMockEnable(enable bool) {
	c.mockEnable = enable
}
