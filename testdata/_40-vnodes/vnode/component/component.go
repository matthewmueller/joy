package component

import "github.com/matthewmueller/golly/testdata/40-vnodes/vnode"

// enforce VNode type
var _ vnode.VNode = (*component)(nil)
var _ Component = (*component)(nil)

// Component interface
type Component interface {
	Render() vnode.VNode
}

// func defaultProps interface{
// 	DefaultProps()
// }

// New fn
func New(com Component) vnode.VNode {
	return &component{
		component: com,
	}
}

// component struct
type component struct {
	component Component
}

// Render fn
func (c *component) Render() vnode.VNode {
	return c.component.Render()
}

// Name fn
func (c *component) Name() string {
	return "#component"
}

// Type fn
func (c *component) Type() int {
	return 0
}

// String fn
func (c *component) String() string {
	willMount, ok := c.component.(componentWillMount)
	if ok {
		willMount.ComponentWillMount()
	}
	return c.Render().String()
}

// shouldComponentUpdate gets invoked before rendering when new props or state are being received.
// This is not called for the initial render or when forceUpdate is used.
type shouldComponentUpdate interface {
	ShouldComponentUpdate(next Component) bool
}

// componentWillUpdate gets invoked immediately before rendering when new props or state are being received.
// This is not called for the initial render.
type componentWillUpdate interface {
	ComponentWillUpdate(next Component)
}

// componentWillReceiveProps gets invoked when a component is receiving new props.
// This method is not called for the initial render.
type componentWillReceiveProps interface {
	ComponentWillReceiveProps(next Component)
}

// componentDidUpdate gets invoked immediately after the component's updates are flushed to the DOM.
// This method is not called for the initial render.
type componentDidUpdate interface {
	ComponentDidUpdate(prev Component)
}

// componentWillMount get invoked once, both on the client and server, immediately before the initial rendering occurs.
type componentWillMount interface {
	ComponentWillMount()
}

// componentWillUnmount gets invoked immediately before a component is unmounted from the DOM.
type componentWillUnmount interface {
	ComponentWillUnmount()
}

// componentDidMount gets invoked once, only on the client (not on the server),
// immediately after the initial rendering occurs.
type componentDidMount interface {
	ComponentDidMount()
}
