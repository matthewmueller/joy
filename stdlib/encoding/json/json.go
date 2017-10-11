package json

import "github.com/matthewmueller/golly/js"

// Marshal a struct into JSON
func Marshal(v interface{}) ([]byte, error) {
	js.Rewrite(`(function(v) {
  try { return [ JSON.stringify(v), null ]  }
  catch (e) { return [ null, e ] }
})($1)`, v)
	return nil, nil
}

// Unmarshal a struct into JSON
func Unmarshal(data []byte, v interface{}) error {

	return nil
}
