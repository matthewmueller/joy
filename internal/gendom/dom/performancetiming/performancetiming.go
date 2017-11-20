package performancetiming

import "github.com/matthewmueller/golly/js"

type PerformanceTiming struct {
}

func (*PerformanceTiming) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PerformanceTiming) GetConnectEnd() (connectEnd uint64) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

func (*PerformanceTiming) GetConnectStart() (connectStart uint64) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

func (*PerformanceTiming) GetDomainLookupEnd() (domainLookupEnd uint64) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceTiming) GetDomainLookupStart() (domainLookupStart uint64) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceTiming) GetDomComplete() (domComplete uint64) {
	js.Rewrite("$<.domComplete")
	return domComplete
}

func (*PerformanceTiming) GetDomContentLoadedEventEnd() (domContentLoadedEventEnd uint64) {
	js.Rewrite("$<.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

func (*PerformanceTiming) GetDomContentLoadedEventStart() (domContentLoadedEventStart uint64) {
	js.Rewrite("$<.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

func (*PerformanceTiming) GetDomInteractive() (domInteractive uint64) {
	js.Rewrite("$<.domInteractive")
	return domInteractive
}

func (*PerformanceTiming) GetDomLoading() (domLoading uint64) {
	js.Rewrite("$<.domLoading")
	return domLoading
}

func (*PerformanceTiming) GetFetchStart() (fetchStart uint64) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

func (*PerformanceTiming) GetLoadEventEnd() (loadEventEnd uint64) {
	js.Rewrite("$<.loadEventEnd")
	return loadEventEnd
}

func (*PerformanceTiming) GetLoadEventStart() (loadEventStart uint64) {
	js.Rewrite("$<.loadEventStart")
	return loadEventStart
}

func (*PerformanceTiming) GetMsFirstPaint() (msFirstPaint uint64) {
	js.Rewrite("$<.msFirstPaint")
	return msFirstPaint
}

func (*PerformanceTiming) GetNavigationStart() (navigationStart uint64) {
	js.Rewrite("$<.navigationStart")
	return navigationStart
}

func (*PerformanceTiming) GetRedirectEnd() (redirectEnd uint64) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

func (*PerformanceTiming) GetRedirectStart() (redirectStart uint64) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

func (*PerformanceTiming) GetRequestStart() (requestStart uint64) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

func (*PerformanceTiming) GetResponseEnd() (responseEnd uint64) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

func (*PerformanceTiming) GetResponseStart() (responseStart uint64) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

func (*PerformanceTiming) GetUnloadEventEnd() (unloadEventEnd uint64) {
	js.Rewrite("$<.unloadEventEnd")
	return unloadEventEnd
}

func (*PerformanceTiming) GetUnloadEventStart() (unloadEventStart uint64) {
	js.Rewrite("$<.unloadEventStart")
	return unloadEventStart
}
