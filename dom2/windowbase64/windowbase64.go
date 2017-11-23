package windowbase64

// js:"WindowBase64,omit"
type WindowBase64 interface {

	// Atob
	Atob(encodedString string) (s string)

	// Btoa
	Btoa(rawString string) (s string)
}
