package window

// js:"GlobalEventHandlers,omit"
type GlobalEventHandlers interface {

	// Onpointercancel
	Onpointercancel() (onpointercancel func(Event))

	// Onpointercancel
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown
	Onpointerdown() (onpointerdown func(Event))

	// Onpointerdown
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter
	Onpointerenter() (onpointerenter func(Event))

	// Onpointerenter
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave
	Onpointerleave() (onpointerleave func(Event))

	// Onpointerleave
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove
	Onpointermove() (onpointermove func(Event))

	// Onpointermove
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout
	Onpointerout() (onpointerout func(Event))

	// Onpointerout
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover
	Onpointerover() (onpointerover func(Event))

	// Onpointerover
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup
	Onpointerup() (onpointerup func(Event))

	// Onpointerup
	SetOnpointerup(onpointerup func(Event))

	// Onwheel
	Onwheel() (onwheel func(Event))

	// Onwheel
	SetOnwheel(onwheel func(Event))
}
