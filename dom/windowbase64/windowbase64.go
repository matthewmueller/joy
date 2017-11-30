package windowbase64

// WindowBase64 interface
// js:"WindowBase64"
type WindowBase64 interface {

	// Atob
	// js:"atob"
	Atob(encodedString string) (s string)

	// Btoa
	// js:"btoa"
	Btoa(rawString string) (s string)
}
