package idb

type KeyboardEventInit struct {
	*EventModifierInit

	key      *string
	location *uint
	repeat   *bool
}
