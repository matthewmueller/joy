package window

import "github.com/matthewmueller/golly/js"

// js:"UIEvent,omit"
type UIEvent interface {
	Event

	// InitUIEvent
	// js:"initUIEvent,rewrite=inituievent"
	InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int)

	// Detail prop
	// js:"detail,rewrite=detail"
	Detail() (detail int)

	// View prop
	// js:"view,rewrite=view"
	View() (view *Window)
}

// inituievent fn
func inituievent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	js.Rewrite("$<.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

// detail prop
func detail() (detail int) {
	js.Rewrite("$<.detail")
	return detail
}

// view prop
func view() (view *Window) {
	js.Rewrite("$<.view")
	return view
}
