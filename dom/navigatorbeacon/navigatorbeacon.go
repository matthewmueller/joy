package navigatorbeacon

// NavigatorBeacon interface
// js:"NavigatorBeacon"
type NavigatorBeacon interface {

	// SendBeacon
	// js:"sendBeacon"
	SendBeacon(url string, data *interface{}) (b bool)
}
