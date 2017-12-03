package main

import "github.com/matthewmueller/joy/macro"

// definition
func definition(a string, b int) ([]string, error) {
	var arr []string
	var err error
	arr = append(arr, "5")
	macro.Raw(`
		var c = String(parseInt(a) + b)
		arr.push(c)
		`)
	arr = append(arr, "9")
	return arr, err
}

func main() {
	arr, err := definition("5", 3)
	if err != nil {
		println(err)
		return
	}
	println(arr)
}
