package window

// GlobalEventHandlers interface
// js:"GlobalEventHandlers"
type GlobalEventHandlers interface {

	// Onpointercancel prop
	// js:"onpointercancel"
	// jsrewrite:"$_.onpointercancel"
	Onpointercancel() (onpointercancel func(Event))

	// SetOnpointercancel prop
	// js:"onpointercancel"
	// jsrewrite:"$_.onpointercancel = $1"
	SetOnpointercancel(onpointercancel func(Event))

	// Onpointerdown prop
	// js:"onpointerdown"
	// jsrewrite:"$_.onpointerdown"
	Onpointerdown() (onpointerdown func(Event))

	// SetOnpointerdown prop
	// js:"onpointerdown"
	// jsrewrite:"$_.onpointerdown = $1"
	SetOnpointerdown(onpointerdown func(Event))

	// Onpointerenter prop
	// js:"onpointerenter"
	// jsrewrite:"$_.onpointerenter"
	Onpointerenter() (onpointerenter func(Event))

	// SetOnpointerenter prop
	// js:"onpointerenter"
	// jsrewrite:"$_.onpointerenter = $1"
	SetOnpointerenter(onpointerenter func(Event))

	// Onpointerleave prop
	// js:"onpointerleave"
	// jsrewrite:"$_.onpointerleave"
	Onpointerleave() (onpointerleave func(Event))

	// SetOnpointerleave prop
	// js:"onpointerleave"
	// jsrewrite:"$_.onpointerleave = $1"
	SetOnpointerleave(onpointerleave func(Event))

	// Onpointermove prop
	// js:"onpointermove"
	// jsrewrite:"$_.onpointermove"
	Onpointermove() (onpointermove func(Event))

	// SetOnpointermove prop
	// js:"onpointermove"
	// jsrewrite:"$_.onpointermove = $1"
	SetOnpointermove(onpointermove func(Event))

	// Onpointerout prop
	// js:"onpointerout"
	// jsrewrite:"$_.onpointerout"
	Onpointerout() (onpointerout func(Event))

	// SetOnpointerout prop
	// js:"onpointerout"
	// jsrewrite:"$_.onpointerout = $1"
	SetOnpointerout(onpointerout func(Event))

	// Onpointerover prop
	// js:"onpointerover"
	// jsrewrite:"$_.onpointerover"
	Onpointerover() (onpointerover func(Event))

	// SetOnpointerover prop
	// js:"onpointerover"
	// jsrewrite:"$_.onpointerover = $1"
	SetOnpointerover(onpointerover func(Event))

	// Onpointerup prop
	// js:"onpointerup"
	// jsrewrite:"$_.onpointerup"
	Onpointerup() (onpointerup func(Event))

	// SetOnpointerup prop
	// js:"onpointerup"
	// jsrewrite:"$_.onpointerup = $1"
	SetOnpointerup(onpointerup func(Event))

	// Onwheel prop
	// js:"onwheel"
	// jsrewrite:"$_.onwheel"
	Onwheel() (onwheel func(Event))

	// SetOnwheel prop
	// js:"onwheel"
	// jsrewrite:"$_.onwheel = $1"
	SetOnwheel(onwheel func(Event))
}
