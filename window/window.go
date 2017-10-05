package window

import "github.com/matthewmueller/golly/js"

// XMLHttpRequest struct
type XMLHttpRequest struct {
	WithCredentials bool
	Onload          func()
	Onerror         func()

	//
	Ok         bool
	Status     int
	StatusText string
	URL        string
}

// Open the request
// js:"open,test=1"
func (x *XMLHttpRequest) Open(method, url string) {
	// return js.Call(x)
}

// Send a payload
// js:"send,test=1"
func (x *XMLHttpRequest) Send(payload []byte) {

}

// GetAllResponseHeaders fn
// js:"getAllResponseHeaders"
func (x *XMLHttpRequest) GetAllResponseHeaders() string {
	return ""
}

// FetchOptions struct
type FetchOptions struct {
	Method      string
	Credentials string
	Headers     map[string]string
	Body        []byte
}

// Fetch a URL
func Fetch(url string, options FetchOptions) (*Response, error) {
	js.Raw(`return window.fetch(url, options)`)

	// request := XMLHttpRequest{}

	// // method
	// method := options.Method
	// if method == "" {
	// 	method = "get"
	// }

	// withCredentials := false
	// if options.Credentials == "include" {
	// 	withCredentials = true
	// }

	// request.Open(method, url)
	// request.WithCredentials = withCredentials

	// request.Onload = func() {

	// }

	// request.Onerror = func() {

	// }

	// request.Send(options.Body)

	return nil, nil
}

// Response struct
type Response struct {
}

// JSON response
func (r *Response) JSON() (string, error) {
	return "", nil
}
