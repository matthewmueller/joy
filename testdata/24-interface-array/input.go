package main

type document struct {
	Fonts []interface{}
}

func main() {
	doc := document{}
	println(doc.Fonts)
}
