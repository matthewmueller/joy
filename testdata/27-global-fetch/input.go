package main

import (
	"github.com/matthewmueller/golly/testdata/27-global-fetch/promise"
)

// Fetch fn
// js:"fetch,global"
func Fetch(url string, options *map[string]string) *promise.Promise {
	return &promise.Promise{}
}

// Response struct
// js:"response,global"
type Response struct {
	responseText string
}

// JSON gets the json
// js:"json"
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
