package perfwidgetexternal

import "github.com/matthewmueller/golly/js"

// PerfWidgetExternal struct
// js:"PerfWidgetExternal,omit"
type PerfWidgetExternal struct {
}

// AddEventListener fn
func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	js.Rewrite("$<.addEventListener($1, $2)", eventType, callback)
}

// GetMemoryUsage fn
func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	js.Rewrite("$<.getMemoryUsage()")
	return u
}

// GetProcessCPUUsage fn
func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	js.Rewrite("$<.getProcessCpuUsage()")
	return u
}

// GetRecentCPUUsage fn
func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentCpuUsage($1)", last)
	return i
}

// GetRecentFrames fn
func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentFrames($1)", last)
	return i
}

// GetRecentMemoryUsage fn
func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentMemoryUsage($1)", last)
	return i
}

// GetRecentPaintRequests fn
func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentPaintRequests($1)", last)
	return i
}

// RemoveEventListener fn
func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	js.Rewrite("$<.removeEventListener($1, $2)", eventType, callback)
}

// RepositionWindow fn
func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	js.Rewrite("$<.repositionWindow($1, $2)", x, y)
}

// ResizeWindow fn
func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	js.Rewrite("$<.resizeWindow($1, $2)", width, height)
}

// ActiveNetworkRequestCount prop
func (*PerfWidgetExternal) ActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	js.Rewrite("$<.activeNetworkRequestCount")
	return activeNetworkRequestCount
}

// AverageFrameTime prop
func (*PerfWidgetExternal) AverageFrameTime() (averageFrameTime float32) {
	js.Rewrite("$<.averageFrameTime")
	return averageFrameTime
}

// AveragePaintTime prop
func (*PerfWidgetExternal) AveragePaintTime() (averagePaintTime float32) {
	js.Rewrite("$<.averagePaintTime")
	return averagePaintTime
}

// ExtraInformationEnabled prop
func (*PerfWidgetExternal) ExtraInformationEnabled() (extraInformationEnabled bool) {
	js.Rewrite("$<.extraInformationEnabled")
	return extraInformationEnabled
}

// IndependentRenderingEnabled prop
func (*PerfWidgetExternal) IndependentRenderingEnabled() (independentRenderingEnabled bool) {
	js.Rewrite("$<.independentRenderingEnabled")
	return independentRenderingEnabled
}

// IrDisablingContentString prop
func (*PerfWidgetExternal) IrDisablingContentString() (irDisablingContentString string) {
	js.Rewrite("$<.irDisablingContentString")
	return irDisablingContentString
}

// IrStatusAvailable prop
func (*PerfWidgetExternal) IrStatusAvailable() (irStatusAvailable bool) {
	js.Rewrite("$<.irStatusAvailable")
	return irStatusAvailable
}

// MaxCPUSpeed prop
func (*PerfWidgetExternal) MaxCPUSpeed() (maxCpuSpeed uint) {
	js.Rewrite("$<.maxCpuSpeed")
	return maxCpuSpeed
}

// PaintRequestsPerSecond prop
func (*PerfWidgetExternal) PaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	js.Rewrite("$<.paintRequestsPerSecond")
	return paintRequestsPerSecond
}

// PerformanceCounter prop
func (*PerfWidgetExternal) PerformanceCounter() (performanceCounter uint64) {
	js.Rewrite("$<.performanceCounter")
	return performanceCounter
}

// PerformanceCounterFrequency prop
func (*PerfWidgetExternal) PerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	js.Rewrite("$<.performanceCounterFrequency")
	return performanceCounterFrequency
}
