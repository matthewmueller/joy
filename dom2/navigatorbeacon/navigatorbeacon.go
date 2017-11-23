package navigatorbeacon

// js:"NavigatorBeacon,omit"
type NavigatorBeacon interface {

	// SendBeacon
	SendBeacon(url string, data *interface{}) (b bool)
}
