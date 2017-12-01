package windowbase64

// WindowBase64 interface
// js:"WindowBase64"
type WindowBase64 interface {

	// Atob
	// js:"atob"
	// jsrewrite:"$_.atob($1)"
	Atob(encodedString string) (s string)

	// Btoa
	// js:"btoa"
	// jsrewrite:"$_.btoa($1)"
	Btoa(rawString string) (s string)
}
