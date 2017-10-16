package vnode

// VNode interface
type VNode interface {
	String() string
	Name() string
	Type() int
}
