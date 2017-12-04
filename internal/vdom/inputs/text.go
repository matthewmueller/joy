package vdom

import "github.com/matthewmueller/joy/macro"

// Text struct
// js:"text,omit"
type Text struct {
	Component
	Node

	value string
}

// S creates a string
func S(value string) *Text {
	macro.Rewrite("$1", value)
	return &Text{
		value: value,
	}
}

// Render returns itself
func (s *Text) Render() Node {
	return s
}

// String returns the text node value
func (s *Text) String() string {
	return s.value
}
