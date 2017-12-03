package json

import "github.com/matthewmueller/joy/macro"

// Marshal a struct into JSON
func Marshal(v interface{}) ([]byte, error) {
	macro.Rewrite(`(function(v) {
  try { return [ JSON.stringify(v), null ]  }
  catch (e) { return [ null, e ] }
})($1)`, v)
	return nil, nil
}

// Unmarshal a struct into JSON
// TODO: create & use object.assign runtime
func Unmarshal(data []byte, v interface{}) error {
	macro.Rewrite(`(function(data, v) {
  try {
    var o = JSON.parse(data)
    for (var k in o) v[k] = o[k]
    return null
  } catch (e) { return e }
})($1, $2)`, data, v)
	return nil
}
