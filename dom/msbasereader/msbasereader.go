package msbasereader

import "github.com/matthewmueller/golly/dom/window"

// MSBaseReader interface
// js:"MSBaseReader"
type MSBaseReader interface {

	// Abort
	// js:"abort"
	// jsrewrite:"$_.abort()"
	Abort()

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

	// ReadyState prop
	// js:"readyState"
	// jsrewrite:"$_.readyState"
	ReadyState() (readyState uint8)

	// Result prop
	// js:"result"
	// jsrewrite:"$_.result"
	Result() (result interface{})
}
