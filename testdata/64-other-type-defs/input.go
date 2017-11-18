package main

// String type
type String string

// func (s String) concat(s2 string) string {
// 	return string(s) + s2
// }

// // String2 type
// type String2 dep.String

// // Map type
// type Map map[string]string

// S fn
func S(s String) String {
	return s
}

func main() {
	println(String("hi"))
	// println(string(4))
	// println(String(4))
	println(S("hi"))

	// m := Map{}
	// m["hi"] = "world"
	// println(m["hi"])
}
