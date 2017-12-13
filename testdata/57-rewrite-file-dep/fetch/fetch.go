package fetch

import (
	"github.com/matthewmueller/joy/macro"
)

// Options struct
type Options struct {
	Method string
}

// Get fn
func Get(url string) (Response, error) {
	macro.Rewrite(`await (async function (unfetch, url) {
		try {
			var res = await unfetch(url)
			return [ res, null ]
		} catch (err) {
			return [ null, err ]
		}
	})($1, $2)
	`, macro.File("./unfetch.js"), url)
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
	macro.Rewrite(`await (async function ($obj) {
		try {
			var $o = await $_.json()
			for (var $k in $o) $obj[$k] = $o[$k]
			return null
		} catch ($e) {
			return $e 
		}
	})($1)`, v)
	return nil
}
