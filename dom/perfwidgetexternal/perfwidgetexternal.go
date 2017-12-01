package perfwidgetexternal

import "github.com/matthewmueller/golly/js"

// PerfWidgetExternal struct
// js:"PerfWidgetExternal,omit"
type PerfWidgetExternal struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	js.Rewrite("$_.addEventListener($1, $2)", eventType, callback)
}

// GetMemoryUsage fn
// js:"getMemoryUsage"
func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	js.Rewrite("$_.getMemoryUsage()")
	return u
}

// GetProcessCPUUsage fn
// js:"getProcessCpuUsage"
func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	js.Rewrite("$_.getProcessCpuUsage()")
	return u
}

// GetRecentCPUUsage fn
// js:"getRecentCpuUsage"
func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	js.Rewrite("$_.getRecentCpuUsage($1)", last)
	return i
}

// GetRecentFrames fn
// js:"getRecentFrames"
func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	js.Rewrite("$_.getRecentFrames($1)", last)
	return i
}

// GetRecentMemoryUsage fn
// js:"getRecentMemoryUsage"
func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	js.Rewrite("$_.getRecentMemoryUsage($1)", last)
	return i
}

// GetRecentPaintRequests fn
// js:"getRecentPaintRequests"
func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	js.Rewrite("$_.getRecentPaintRequests($1)", last)
	return i
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	js.Rewrite("$_.removeEventListener($1, $2)", eventType, callback)
}

// RepositionWindow fn
// js:"repositionWindow"
func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	js.Rewrite("$_.repositionWindow($1, $2)", x, y)
}

// ResizeWindow fn
// js:"resizeWindow"
func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	js.Rewrite("$_.resizeWindow($1, $2)", width, height)
}

// ActiveNetworkRequestCount prop
// js:"activeNetworkRequestCount"
func (*PerfWidgetExternal) ActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	js.Rewrite("$_.activeNetworkRequestCount")
	return activeNetworkRequestCount
}

// AverageFrameTime prop
// js:"averageFrameTime"
func (*PerfWidgetExternal) AverageFrameTime() (averageFrameTime float32) {
	js.Rewrite("$_.averageFrameTime")
	return averageFrameTime
}

// AveragePaintTime prop
// js:"averagePaintTime"
func (*PerfWidgetExternal) AveragePaintTime() (averagePaintTime float32) {
	js.Rewrite("$_.averagePaintTime")
	return averagePaintTime
}

// ExtraInformationEnabled prop
// js:"extraInformationEnabled"
func (*PerfWidgetExternal) ExtraInformationEnabled() (extraInformationEnabled bool) {
	js.Rewrite("$_.extraInformationEnabled")
	return extraInformationEnabled
}

// IndependentRenderingEnabled prop
// js:"independentRenderingEnabled"
func (*PerfWidgetExternal) IndependentRenderingEnabled() (independentRenderingEnabled bool) {
	js.Rewrite("$_.independentRenderingEnabled")
	return independentRenderingEnabled
}

// IrDisablingContentString prop
// js:"irDisablingContentString"
func (*PerfWidgetExternal) IrDisablingContentString() (irDisablingContentString string) {
	js.Rewrite("$_.irDisablingContentString")
	return irDisablingContentString
}

// IrStatusAvailable prop
// js:"irStatusAvailable"
func (*PerfWidgetExternal) IrStatusAvailable() (irStatusAvailable bool) {
	js.Rewrite("$_.irStatusAvailable")
	return irStatusAvailable
}

// MaxCPUSpeed prop
// js:"maxCpuSpeed"
func (*PerfWidgetExternal) MaxCPUSpeed() (maxCpuSpeed uint) {
	js.Rewrite("$_.maxCpuSpeed")
	return maxCpuSpeed
}

// PaintRequestsPerSecond prop
// js:"paintRequestsPerSecond"
func (*PerfWidgetExternal) PaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	js.Rewrite("$_.paintRequestsPerSecond")
	return paintRequestsPerSecond
}

// PerformanceCounter prop
// js:"performanceCounter"
func (*PerfWidgetExternal) PerformanceCounter() (performanceCounter uint64) {
	js.Rewrite("$_.performanceCounter")
	return performanceCounter
}

// PerformanceCounterFrequency prop
// js:"performanceCounterFrequency"
func (*PerfWidgetExternal) PerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	js.Rewrite("$_.performanceCounterFrequency")
	return performanceCounterFrequency
}
