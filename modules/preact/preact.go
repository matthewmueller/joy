package preact

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/dom"
)

// IProps interface
type IProps interface{}

// IState interface
type IState interface{}

// IContext interface
type IContext interface{}

// VNode struct
type VNode struct {
	NodeType   int
	NodeName   string
	NodeValue  string
	Component  IComponent
	Children   []VNode
	Attributes map[string]interface{}
	Key        *string
}

// IComponent interfaces
type IComponent interface {
	Render(props IProps, state IState) *VNode
}

// C creates a VNode from a Component
func C(component IComponent, props map[string]interface{}, children ...VNode) VNode {
	return VNode{
		NodeType:   0,
		Component:  component,
		Attributes: props,
		Children:   children,
	}
}

// H creates a VNode from an HTML element
func H(tag string, attrs map[string]interface{}, children ...VNode) VNode {
	return VNode{
		NodeType:   1,
		NodeName:   tag,
		Attributes: attrs,
		Children:   children,
	}
}

// T creates a textnode
func T(text string) VNode {
	return VNode{
		NodeType:  3,
		NodeValue: text,
	}
}

// Render a function
func Render(vnode VNode, parent *dom.Node, merge *dom.Node) *dom.Node {
	return diff(merge, vnode, nil, false, parent, false)
}

/** Queue of components that have been mounted and are awaiting componentDidMount */
var mounts = []IComponent{}

/** Diff recursion count, used to track the end of the diff cycle. */
var diffLevel = 0

/** Global flag indicating if the diff is currently within an SVG */
var isSvgMode = false

/** Global flag indicating if the diff is performing hydration */
var hydrating = false

func diff(
	dom *dom.Node,
	vnode VNode,
	context IContext,
	mountAll bool,
	parent *dom.Node,
	componentRoot bool,
) *dom.Node {

	log.Infof("got all?")
	return nil
}

// New fn
// func New(component Component) Preact {
// 	return Preact{
// 		component: component,
// 	}
// }

// // Component fn
// type Component interface {
// 	Render() *VNode
// }

// // Preact is a component wrapped in the preact library
// type Preact struct {
// 	component Component
// }

// // Render fn
// // preact.render(<Page class="hi" />, document.body, document.body.firstChild)
// func Render(vnode VNode, root document.Node) {
// 	root.AppendChild(document.CreateElement("hi world!"))
// }

// // Div creates a div
// func Div(text string) VNode {
// 	return VNode{
// 		NodeName:   "div",
// 		Attributes: nil,
// 		Children: []VNode{
// 			VNode{
// 				NodeName:  "#text",
// 				NodeValue: text,
// 			},
// 		},
// 	}
// }

// // // Cops holds COntext, Props and State received in the lifecycle methods.
// // // Note that any of these can be nil, depending on the context.
// // // type Cops struct {
// // // 	Context Context
// // // 	Props   Props
// // // 	State   State
// // // }

// // // StateInitializer sets up the initial state.
// // type StateInitializer interface {
// // 	GetInitialState() State
// // }

// // // ChildContextProvider provides the context for the children.
// // //
// // // The GetChildContext function will be called when the state or props changes.
// // // In order to update data in the context, trigger a local state update with this.SetState.
// // // This will trigger a new context and changes will be received by the children.
// // //
// // // GetChildContext will also be called once in the init phase, to determine the types for
// // // the context properties. The this will be nil in this single invocation, and there is no need to return
// // // real data as long as the types are real (in cases where this is an expensive operation).
// // //
// // // See https://facebook.github.io/react/docs/context.html
// // type ChildContextProvider interface {
// // 	GetChildContext() Context
// // }

// // // ShouldComponentUpdate gets invoked before rendering when new props or state are being received.
// // // This is not called for the initial render or when forceUpdate is used.
// // type ShouldComponentUpdate interface {
// // 	ShouldComponentUpdate(next Cops) bool
// // }

// // // ComponentWillUpdate gets invoked immediately before rendering when new props or state are being received.
// // // This is not called for the initial render.
// // type ComponentWillUpdate interface {
// // 	ComponentWillUpdate(next Cops)
// // }

// // // ComponentWillReceiveProps gets invoked when a component is receiving new props.
// // // This method is not called for the initial render.
// // type ComponentWillReceiveProps interface {
// // 	ComponentWillReceiveProps(next Cops)
// // }

// // // ComponentDidUpdate gets invoked immediately after the component's updates are flushed to the DOM.
// // // This method is not called for the initial render.
// // type ComponentDidUpdate interface {
// // 	ComponentDidUpdate(prev Cops)
// // }

// // // ComponentWillMount get invoked once, both on the client and server, immediately before the initial rendering occurs.
// // type ComponentWillMount interface {
// // 	ComponentWillMount()
// // }

// // // ComponentWillUnmount gets invoked immediately before a component is unmounted from the DOM.
// // type ComponentWillUnmount interface {
// // 	ComponentWillUnmount()
// // }

// // // ComponentDidMount gets invoked once, only on the client (not on the server),
// // // immediately after the initial rendering occurs.
// // type ComponentDidMount interface {
// // 	ComponentDidMount()
// // }
