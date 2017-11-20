package navigatorbeacon

import "github.com/matthewmueller/golly/js"

type NavigatorBeacon struct {
}

func (*NavigatorBeacon) SendBeacon(url string, data interface{}) (b bool) {
	js.Rewrite("$<.sendBeacon($1, $2)", url, data)
	return b
}
