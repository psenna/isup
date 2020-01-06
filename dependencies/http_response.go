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
	Error         string
}