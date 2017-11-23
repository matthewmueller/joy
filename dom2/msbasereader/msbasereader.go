package msbasereader

import "github.com/matthewmueller/golly/dom2/window"

// js:"MSBaseReader,omit"
type MSBaseReader interface {

	// Abort
	Abort()

	// Onabort
	Onabort() (onabort func(window.Event))

	// Onabort
	SetOnabort(onabort func(window.Event))

	// Onerror
	Onerror() (onerror func(window.Event))

	// Onerror
	SetOnerror(onerror func(window.Event))

	// Onload
	Onload() (onload func(window.Event))

	// Onload
	SetOnload(onload func(window.Event))

	// Onloadend
	Onloadend() (onloadend func(window.Event))

	// Onloadend
	SetOnloadend(onloadend func(window.Event))

	// Onloadstart
	Onloadstart() (onloadstart func(window.Event))

	// Onloadstart
	SetOnloadstart(onloadstart func(window.Event))

	// Onprogress
	Onprogress() (onprogress func(window.Event))

	// Onprogress
	SetOnprogress(onprogress func(window.Event))

	// ReadyState
	ReadyState() (readyState uint8)

	// Result
	Result() (result interface{})
}
