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

// Github struct
// js:"github,global"
type Github struct {
	Login string `json:"login,omitempty" js:"login"`
	Name  string `json:"name,omitempty" js:"name"`
	Bio   string `json:"bio,omitempty" js:"bio"`
}

func main() {
	p := Fetch("https://api.github.com/users/matthewmueller", nil)
	ch := make(chan Github, 1)

	p.Then(func(v interface{}) interface{} {
		res := v.(Response)
		return res.JSON()
	}).Then(func(v interface{}) interface{} {
		ch <- v.(Github)
		return nil
	}).Catch(func(e error) interface{} {
		println(e)
		return nil
	})

	gh := <-ch
	println(gh.Login, "-", gh.Name)
}
