package main

// Subthing struct
type subthing struct {
	name string
}

// Name string
func (s *subthing) Name() string {
	return s.name
}

// Thing struct
type Thing struct {
	*subthing
}

func main() {
	s := Thing{&subthing{"matt"}}
	println(s.subthing.Name())
}
