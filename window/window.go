package window

import "encoding/json"

// XMLHttpRequest struct
type XMLHttpRequest struct {
}

// FetchOptions struct
type FetchOptions struct {
	Headers map[string]string
}

// Fetch a URL
func Fetch(url string, options *FetchOptions) (*Response, error) {
	return nil, nil
}

// Response struct
type Response struct {
}

// JSON response
func (r *Response) JSON() (json.RawMessage, error) {
	return nil, nil
}
