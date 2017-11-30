package window

// GlobalEventHandlers interface
// js:"GlobalEventHandlers"
type GlobalEventHandlers interface {

	// Onpointercancel prop
	// js:"onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// SetOnpointercancel prop
	// js:"onpointercancel"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// SetOnpointerdown prop
	// js:"onpointerdown"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// SetOnpointerenter prop
	// js:"onpointerenter"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// SetOnpointerleave prop
	// js:"onpointerleave"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove"
	Onpointermove() (onpointermove func(Event))

	// SetOnpointermove prop
	// js:"onpointermove"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout"
	Onpointerout() (onpointerout func(Event))

	// SetOnpointerout prop
	// js:"onpointerout"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover"
	Onpointerover() (onpointerover func(Event))

	// SetOnpointerover prop
	// js:"onpointerover"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup"
	Onpointerup() (onpointerup func(Event))

	// SetOnpointerup prop
	// js:"onpointerup"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel"
	Onwheel() (onwheel func(Event))

	// SetOnwheel prop
	// js:"onwheel"
	SetOnwheel(onwheel func(Event))
}
