package dependencies

// HTTPRequest A request for a http call
// If InsecureCall is true the ssl certificate is not validated
// TimeOut is the call timeout in milisecounds, default is 10000 ms, set 0 to no timeout
type HTTPRequest struct {
	URL          string
	Method       string
	Headers      string
	Body         string
	InsecureCall bool
	TimeOut      int64
}
