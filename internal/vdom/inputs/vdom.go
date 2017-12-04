package vdom

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

//go:generate go run internal/gen.go

// Use fn
// js:"Use,omit"
func Use(pragma, filepath string) {
}

// Pragma is a reference to how elements get created
// js:"Pragma,omit"
func Pragma() string {
	return ""
}

// File is a reference to the vdom library itself
// js:"File,omit"
func File() string {
	return ""
}

// Component struct
type Component interface {
	// js:"render"
	Render() Node
	// js:"setState"
	SetState(state interface{})
	// js:"forceUpdate"
	ForceUpdate()
}

// Child interface
type Child interface {
	Render() Node
}

// Node interface
type Node interface {
	String() string
}

// Render the component
func Render(component Child, parent window.Node, merge window.Node) {
	macro.Rewrite("$1.render($2, $3, $4)", File(), component, parent, merge)
}

// String turns the component into a string
func String(component Child) string {
	return component.Render().String()
}
