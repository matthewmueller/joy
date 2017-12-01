package navigatoronline

// NavigatorOnLine interface
// js:"NavigatorOnLine"
type NavigatorOnLine interface {

	// OnLine prop
	// js:"onLine"
	// jsrewrite:"$_.onLine"
	OnLine() (onLine bool)
}
