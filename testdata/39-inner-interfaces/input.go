package main

// Stringer interface
type Stringer interface {
	String() string
}

type stringer struct {
	value string
}

func (s *stringer) String() string {
	return s.value
}

// Subthing struct
type Subthing struct {
	first Stringer
	last  Stringer
}

func (s *Subthing) firstName() Stringer {
	return s.first
}

func (s *Subthing) lastName() Stringer {
	return s.last
}

// Name string
func (s *Subthing) Name() string {
	return s.firstName().String() + " " + s.lastName().String()
}

// Thing struct
type Thing struct {
	*Subthing
}

func main() {
	s := Thing{&Subthing{&stringer{"matt"}, &stringer{"mueller"}}}
	println(s.Subthing.Name())
}
