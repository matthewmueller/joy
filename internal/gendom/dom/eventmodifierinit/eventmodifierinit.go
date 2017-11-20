package eventmodifierinit

type EventModifierInit struct {
	*UIEventInit

	altKey             *bool
	ctrlKey            *bool
	metaKey            *bool
	modifierAltGraph   *bool
	modifierCapsLock   *bool
	modifierFn         *bool
	modifierFnLock     *bool
	modifierHyper      *bool
	modifierNumLock    *bool
	modifierOS         *bool
	modifierScrollLock *bool
	modifierSuper      *bool
	modifierSymbol     *bool
	modifierSymbolLock *bool
	shiftKey           *bool
}
