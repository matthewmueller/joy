package main

type component interface {
	// js:"render"
	Render() string
}

type page struct {
}

// js:"render"
func (*page) Render() string {
	return "page!"
}

func main() {
	p := newpage()
	println(p.Render())

	i := newiface()
	println(i.Render())
}

func newpage() *page {
	return &page{}
}

func newiface() component {
	return &page{}
}
