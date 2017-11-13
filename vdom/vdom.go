package vdom

//go:generate go run internal/gen.go

// Use fn
// js:"use,omit"
func Use(pragma, filepath string) {
}

// Pragma is a reference to how elements get created
// js:"pragma,omit"
func Pragma() string {
	return ""
}

// File is a reference to the vdom library itself
// js:"file,omit"
func File() string {
	return ""
}

// Component struct
type Component interface {
	Render() Node
	SetState(state interface{})
}

// Child interface
type Child interface {
	Render() Node
}

// // SetState struct
// type SetState struct{}

// // SetState struct
// // js:"setState,omit"
// func (s *SetState) SetState(state interface{}) {
// 	return
// }

// type Component interface {
// 	Render() Node
// }

// SetState interface
// type SetState interface {
// 	SetState(state interface{})
// }

// Node interface
type Node interface {
	String() string
}
