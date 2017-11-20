package perfwidgetexternal

import "github.com/matthewmueller/golly/js"

type PerfWidgetExternal struct {
}

func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	js.Rewrite("$<.addEventListener($1, $2)", eventType, callback)
}

func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	js.Rewrite("$<.getMemoryUsage()")
	return u
}

func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	js.Rewrite("$<.getProcessCpuUsage()")
	return u
}

func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentCpuUsage($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentFrames($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentMemoryUsage($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentPaintRequests($1)", last)
	return i
}

func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	js.Rewrite("$<.removeEventListener($1, $2)", eventType, callback)
}

func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	js.Rewrite("$<.repositionWindow($1, $2)", x, y)
}

func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	js.Rewrite("$<.resizeWindow($1, $2)", width, height)
}

func (*PerfWidgetExternal) GetActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	js.Rewrite("$<.activeNetworkRequestCount")
	return activeNetworkRequestCount
}

func (*PerfWidgetExternal) GetAverageFrameTime() (averageFrameTime float32) {
	js.Rewrite("$<.averageFrameTime")
	return averageFrameTime
}

func (*PerfWidgetExternal) GetAveragePaintTime() (averagePaintTime float32) {
	js.Rewrite("$<.averagePaintTime")
	return averagePaintTime
}

func (*PerfWidgetExternal) GetExtraInformationEnabled() (extraInformationEnabled bool) {
	js.Rewrite("$<.extraInformationEnabled")
	return extraInformationEnabled
}

func (*PerfWidgetExternal) GetIndependentRenderingEnabled() (independentRenderingEnabled bool) {
	js.Rewrite("$<.independentRenderingEnabled")
	return independentRenderingEnabled
}

func (*PerfWidgetExternal) GetIrDisablingContentString() (irDisablingContentString string) {
	js.Rewrite("$<.irDisablingContentString")
	return irDisablingContentString
}

func (*PerfWidgetExternal) GetIrStatusAvailable() (irStatusAvailable bool) {
	js.Rewrite("$<.irStatusAvailable")
	return irStatusAvailable
}

func (*PerfWidgetExternal) GetMaxCPUSpeed() (maxCpuSpeed uint) {
	js.Rewrite("$<.maxCpuSpeed")
	return maxCpuSpeed
}

func (*PerfWidgetExternal) GetPaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	js.Rewrite("$<.paintRequestsPerSecond")
	return paintRequestsPerSecond
}

func (*PerfWidgetExternal) GetPerformanceCounter() (performanceCounter uint64) {
	js.Rewrite("$<.performanceCounter")
	return performanceCounter
}

func (*PerfWidgetExternal) GetPerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	js.Rewrite("$<.performanceCounterFrequency")
	return performanceCounterFrequency
}
