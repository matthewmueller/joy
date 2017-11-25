// +build macro

package main

import "github.com/matthewmueller/golly/js"

func joyNodeName() string {
	js.Rewrite("$<.nodeName")
	return ""
}
