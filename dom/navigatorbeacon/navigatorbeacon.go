package navigatorbeacon

// NavigatorBeacon interface
// js:"NavigatorBeacon"
type NavigatorBeacon interface {

	// SendBeacon
	// js:"sendBeacon"
	// jsrewrite:"$_.sendBeacon($1, $2)"
	SendBeacon(url string, data *interface{}) (b bool)
}
