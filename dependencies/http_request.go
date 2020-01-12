package dependencies

import "net/http"

// HTTPRequest A request for a http call
// If InsecureRequest is true the ssl certificate is not validated
// TimeOut is the call timeout in milisecounds, default is 2000 ms, max is 60000
type HTTPRequest struct {
	url             string
	method          string
	headers         map[string]interface{}
	formParams      map[string]interface{}
	insecureRequest bool
	timeOut         int
}

const (
	timeOut    = 2000
	minTimeOut = 50
	maxTimeout = 60000
)

// GetHTTPRequest Instantiate a HTTP request object
func GetHTTPRequest(method string, url string) HTTPRequest {
	h := HTTPRequest{url: url, method: method, timeOut: timeOut}

	return h
}

// SetHeaders Set headers values
func (h HTTPRequest) SetHeaders(headers map[string]interface{}) HTTPRequest {
	if h.headers == nil {
		h.headers = make(map[string]interface{})
	}

	for index, value := range headers {
		h.headers[index] = value
	}
	return h
}

// SetFormParams Set forms values
func (h HTTPRequest) SetFormParams(formParams map[string]interface{}) HTTPRequest {
	if h.formParams == nil {
		h.formParams = make(map[string]interface{})
	}

	for index, value := range formParams {
		h.formParams[index] = value
	}

	return h
}

// SetInsecureRequest Set if request is insecure (no cert validation)
func (h HTTPRequest) SetInsecureRequest(InsecureRequest bool) HTTPRequest {
	h.insecureRequest = InsecureRequest
	return h
}

// SetTimeOut Set request timeout
func (h HTTPRequest) SetTimeOut(timeOut int) HTTPRequest {
	h.timeOut = timeOut

	if h.timeOut < minTimeOut {
		h.timeOut = minTimeOut
	}

	if h.timeOut > maxTimeout {
		h.timeOut = maxTimeout
	}

	return h
}

// GetTimeOut Get request timeout
func (h HTTPRequest) GetTimeOut() int {
	return h.timeOut
}

// GetInsecureRequest Get if request is insecure
func (h HTTPRequest) GetInsecureRequest() bool {
	return h.insecureRequest
}

// ToGoHTTPRequest Create a go http.Request from a HTTPRequest
func (h HTTPRequest) ToGoHTTPRequest() (*http.Request, error) {
	return http.NewRequest(h.method, h.url, nil)
}
