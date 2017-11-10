package fetch

import (
	"github.com/matthewmueller/golly/js"
)

// Options struct
type Options struct {
	Method string
}

// Fetch fn
func Fetch(url string, options *Options) (Response, error) {
	js.Rewrite(`
	try {
		var res = await unfetch(url, options)
		return [ res, null ]
	} catch (err) {
		return [ null, err ]
	}
`, js.RawFile("./unfetch.js"), url, options)
	return Response{}, nil
}

// Get fn
func Get(url string) (Response, error) {
	js.Rewrite(`await (async function (unfetch, url) {
		try {
			var res = await unfetch(url)
			return [ res, null ]
		} catch (err) {
			return [ null, err ]
		}
	})($1, $2)
	`, js.RawFile("./unfetch.js"), url)
	return Response{}, nil
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

// JSON marshals the json
// js:"json,async"
func (r *Response) JSON(v interface{}) error {
	js.Rewrite(`await (async function ($obj) {
		try {
			var $o = await $<.json()
			for (var $k in $o) $obj[$k] = $o[$k]
			return null
		} catch ($e) {
			return $e
		}
	})($1)`, v)
	return nil
}
