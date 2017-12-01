package window

// WindowConsole interface
// js:"WindowConsole"
type WindowConsole interface {

	// Console prop
	// js:"console"
	// jsrewrite:"$_.console"
	Console() (console *Console)
}
