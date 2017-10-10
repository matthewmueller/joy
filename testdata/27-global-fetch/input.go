package main

import (
	"github.com/matthewmueller/golly/testdata/27-global-fetch/promise"
)

// Fetch fn
// js:"global,fetch"
func Fetch(url string, options *map[string]string) *promise.Promise {
	return &promise.Promise{}
}

// Response struct
type Response struct {
	responseText string
}

// JSON gets the json
func (r *Response) JSON() string {
	return r.responseText
}

func main() {
	p := Fetch("http://google.com", nil)

	p.Then(func(v interface{}) interface{} {
		res := v.(Response)
		return res.JSON()
	}).Then(func(v interface{}) interface{} {
		res := v.(string)
		println(res)
		return nil
	}).Catch(func(e error) interface{} {
		println(e)
		return nil
	})
}
