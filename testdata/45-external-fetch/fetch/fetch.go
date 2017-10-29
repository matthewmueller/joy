package fetch

import (
	"github.com/matthewmueller/golly/js"
)

// unfetch
var unfetch = js.RawFile("./unfetch.js")

// Options struct
type Options struct {
	Method string
}

// Fetch fn
func Fetch(url string, options *Options) (*Response, error) {
	js.Raw(`
	try {
		var res = await unfetch(url, options)
		return [ res, null ]
	} catch (err) {
		return [ null, err ]
	}
`, unfetch)
	return nil, nil
}

// Response struct
// js:"response,omit"
type Response struct {
}

// Text gets the textual representation
// js:"text,async"
func (r *Response) Text() []byte {
	return nil
}
