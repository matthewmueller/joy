package fetch

import (
	"github.com/matthewmueller/joy/macro"
)

// unfetch
var unfetch = macro.RawFile("./unfetch.js")

// Options struct
type Options struct {
	Method string
}

// Fetch fn
func Fetch(url string, options *Options) (*Response, error) {
	macro.Raw(`
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
