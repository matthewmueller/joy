package main

import "github.com/matthewmueller/golly/testdata/64-other-type-defs/dep"

// String type
type String string

// Integer type
type Integer int

// func (s String) concat(s2 string) string {
// 	return string(s) + s2
// }

// String2 type
type String2 dep.String

// Map type
type Map map[string]string

// Arr type
type Arr []string

// S fn
func S(s String) String {
	return s
}

func main() {
	println(String("hi"))
	println(String2("hi"))
	println(Integer(5))
	println(S("hi"))

	// println(String("hi").concat("cool"))

	m := Map{"a": "b"}
	m["hello"] = "world"
	println(m["hello"])
	println(m["a"])

	var a Arr
	a = append(a, "array")
	println(a[0])
}
