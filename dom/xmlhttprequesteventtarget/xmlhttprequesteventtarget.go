package xmlhttprequesteventtarget

import "github.com/matthewmueller/joy/dom/window"

// XMLHTTPRequestEventTarget interface
// js:"XMLHTTPRequestEventTarget"
type XMLHTTPRequestEventTarget interface {

	// Onabort prop
	// js:"onabort"
	// jsrewrite:"$_.onabort"
	Onabort() (onabort func(window.Event))

	// SetOnabort prop
	// js:"onabort"
	// jsrewrite:"$_.onabort = $1"
	SetOnabort(onabort func(window.Event))

	// Onerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror"
	Onerror() (onerror func(window.Event))

	// SetOnerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror = $1"
	SetOnerror(onerror func(window.Event))

	// Onload prop
	// js:"onload"
	// jsrewrite:"$_.onload"
	Onload() (onload func(window.Event))

	// SetOnload prop
	// js:"onload"
	// jsrewrite:"$_.onload = $1"
	SetOnload(onload func(window.Event))

	// Onloadend prop
	// js:"onloadend"
	// jsrewrite:"$_.onloadend"
	Onloadend() (onloadend func(window.Event))

	// SetOnloadend prop
	// js:"onloadend"
	// jsrewrite:"$_.onloadend = $1"
	SetOnloadend(onloadend func(window.Event))

	// Onloadstart prop
	// js:"onloadstart"
	// jsrewrite:"$_.onloadstart"
	Onloadstart() (onloadstart func(window.Event))

	// SetOnloadstart prop
	// js:"onloadstart"
	// jsrewrite:"$_.onloadstart = $1"
	SetOnloadstart(onloadstart func(window.Event))

	// Onprogress prop
	// js:"onprogress"
	// jsrewrite:"$_.onprogress"
	Onprogress() (onprogress func(window.Event))

	// SetOnprogress prop
	// js:"onprogress"
	// jsrewrite:"$_.onprogress = $1"
	SetOnprogress(onprogress func(window.Event))

	// Ontimeout prop
	// js:"ontimeout"
	// jsrewrite:"$_.ontimeout"
	Ontimeout() (ontimeout func(window.Event))

	// SetOntimeout prop
	// js:"ontimeout"
	// jsrewrite:"$_.ontimeout = $1"
	SetOntimeout(ontimeout func(window.Event))
}
