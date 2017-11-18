package main

// String type
type String string
type Map map[string]string

func (s String) concat(s2 string) String {
	return s + String(s2)
}

// S fn
func S(s String) String {
	return s
}

func main() {
	println(String("hi").concat("cool"))
	println(string(4))
	println(String(4))
	println(S("hi"))

}
