package element

// Element i
type Element interface {
	Render() string
}

// New fn
func New(tag string, child Element) Element {
	return &element{
		tag:   tag,
		child: child,
	}
}

type element struct {
	tag   string
	child Element
}

func (c *element) Render() string {
	if c.child == nil {
		return "<" + c.tag + "></" + c.tag + ">"
	}
	return "<" + c.tag + ">" + c.child.Render() + "</" + c.tag + ">"
}
