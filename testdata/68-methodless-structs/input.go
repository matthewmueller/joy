package main

type Options struct {
	a string
	b int
	c C
}

type C struct {
	d int
}

func main() {
	o := &Options{a: "hi"}
	println(o.a, o.b, o.c.d)
}
