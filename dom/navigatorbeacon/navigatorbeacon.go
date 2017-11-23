package navigatorbeacon

import "github.com/matthewmueller/golly/js"

// NavigatorBeacon struct
// js:"NavigatorBeacon,omit"
type NavigatorBeacon struct {
}

// SendBeacon fn
func (*NavigatorBeacon) SendBeacon(url string, data *interface{}) (b bool) {
	js.Rewrite("$<.sendBeacon($1, $2)", url, data)
	return b
}
