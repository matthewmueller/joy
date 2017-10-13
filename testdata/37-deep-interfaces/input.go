package main

// VNode interface
type VNode interface {
	String() string
}

// Component interface
type Component interface {
	Render() string
}

type vnode struct {
	c *Component
}

func (v *vnode) String() string {
	return "earth"
}

func newComponent(c Component) VNode {
	return &vnode{&c}
}

type index struct {
	location string
}

func (i *index) Render() string {
	return i.location
}

func newIndex(location string) VNode {
	return newComponent(&index{location})
}

func main() {
	c := newIndex("mars")
	println(c.String())
}
