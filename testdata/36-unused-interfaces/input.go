package main

var _ Component = (*component)(nil)

// Component interface
type Component interface {
	String() string
}

type component struct {
	component string
}

func (c component) String() string {
	return c.component
}

func (c component) another() string {
	return "another"
}

type blah struct {
	Blah string
}

func (b blah) String() string {
	return b.Blah
}

// New component
func New(name string) Component {
	return &component{name}
}

func main() {
	c := New("awesomeness")
	println(c.String())
}
