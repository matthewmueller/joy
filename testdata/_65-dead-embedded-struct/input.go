package main

// a struct
type a struct {
	Age int
}

// B struct
type B struct {
	a

	Name string
}

func main() {
	b := B{Name: "matt"}
	println(b.Name, b.Age)
}
