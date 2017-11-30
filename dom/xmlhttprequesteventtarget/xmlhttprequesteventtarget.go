package xmlhttprequesteventtarget

import "github.com/matthewmueller/golly/dom/window"

// XMLHTTPRequestEventTarget interface
// js:"XMLHTTPRequestEventTarget"
type XMLHTTPRequestEventTarget interface {

	// Onabort prop
	// js:"onabort"
	Onabort() (onabort func(window.Event))

	// SetOnabort prop
	// js:"onabort"
	SetOnabort(onabort func(window.Event))

	// Onerror prop
	// js:"onerror"
	Onerror() (onerror func(window.Event))

	// SetOnerror prop
	// js:"onerror"
	SetOnerror(onerror func(window.Event))

	// Onload prop
	// js:"onload"
	Onload() (onload func(window.Event))

	// SetOnload prop
	// js:"onload"
	SetOnload(onload func(window.Event))

	// Onloadend prop
	// js:"onloadend"
	Onloadend() (onloadend func(window.Event))

	// SetOnloadend prop
	// js:"onloadend"
	SetOnloadend(onloadend func(window.Event))

	// Onloadstart prop
	// js:"onloadstart"
	Onloadstart() (onloadstart func(window.Event))

	// SetOnloadstart prop
	// js:"onloadstart"
	SetOnloadstart(onloadstart func(window.Event))

	// Onprogress prop
	// js:"onprogress"
	Onprogress() (onprogress func(window.Event))

	// SetOnprogress prop
	// js:"onprogress"
	SetOnprogress(onprogress func(window.Event))

	// Ontimeout prop
	// js:"ontimeout"
	Ontimeout() (ontimeout func(window.Event))

	// SetOntimeout prop
	// js:"ontimeout"
	SetOntimeout(ontimeout func(window.Event))
}
