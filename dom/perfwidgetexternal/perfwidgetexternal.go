package perfwidgetexternal

import "github.com/matthewmueller/joy/macro"

// PerfWidgetExternal struct
// js:"PerfWidgetExternal,omit"
type PerfWidgetExternal struct {
}

// AddEventListener fn
// js:"addEventListener"
func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	macro.Rewrite("$_.addEventListener($1, $2)", eventType, callback)
}

// GetMemoryUsage fn
// js:"getMemoryUsage"
func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	macro.Rewrite("$_.getMemoryUsage()")
	return u
}

// GetProcessCPUUsage fn
// js:"getProcessCpuUsage"
func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	macro.Rewrite("$_.getProcessCpuUsage()")
	return u
}

// GetRecentCPUUsage fn
// js:"getRecentCpuUsage"
func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentCpuUsage($1)", last)
	return i
}

// GetRecentFrames fn
// js:"getRecentFrames"
func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentFrames($1)", last)
	return i
}

// GetRecentMemoryUsage fn
// js:"getRecentMemoryUsage"
func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentMemoryUsage($1)", last)
	return i
}

// GetRecentPaintRequests fn
// js:"getRecentPaintRequests"
func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentPaintRequests($1)", last)
	return i
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	macro.Rewrite("$_.removeEventListener($1, $2)", eventType, callback)
}

// RepositionWindow fn
// js:"repositionWindow"
func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	macro.Rewrite("$_.repositionWindow($1, $2)", x, y)
}

// ResizeWindow fn
// js:"resizeWindow"
func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	macro.Rewrite("$_.resizeWindow($1, $2)", width, height)
}

// ActiveNetworkRequestCount prop
// js:"activeNetworkRequestCount"
func (*PerfWidgetExternal) ActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	macro.Rewrite("$_.activeNetworkRequestCount")
	return activeNetworkRequestCount
}

// AverageFrameTime prop
// js:"averageFrameTime"
func (*PerfWidgetExternal) AverageFrameTime() (averageFrameTime float32) {
	macro.Rewrite("$_.averageFrameTime")
	return averageFrameTime
}

// AveragePaintTime prop
// js:"averagePaintTime"
func (*PerfWidgetExternal) AveragePaintTime() (averagePaintTime float32) {
	macro.Rewrite("$_.averagePaintTime")
	return averagePaintTime
}

// ExtraInformationEnabled prop
// js:"extraInformationEnabled"
func (*PerfWidgetExternal) ExtraInformationEnabled() (extraInformationEnabled bool) {
	macro.Rewrite("$_.extraInformationEnabled")
	return extraInformationEnabled
}

// IndependentRenderingEnabled prop
// js:"independentRenderingEnabled"
func (*PerfWidgetExternal) IndependentRenderingEnabled() (independentRenderingEnabled bool) {
	macro.Rewrite("$_.independentRenderingEnabled")
	return independentRenderingEnabled
}

// IrDisablingContentString prop
// js:"irDisablingContentString"
func (*PerfWidgetExternal) IrDisablingContentString() (irDisablingContentString string) {
	macro.Rewrite("$_.irDisablingContentString")
	return irDisablingContentString
}

// IrStatusAvailable prop
// js:"irStatusAvailable"
func (*PerfWidgetExternal) IrStatusAvailable() (irStatusAvailable bool) {
	macro.Rewrite("$_.irStatusAvailable")
	return irStatusAvailable
}

// MaxCPUSpeed prop
// js:"maxCpuSpeed"
func (*PerfWidgetExternal) MaxCPUSpeed() (maxCpuSpeed uint) {
	macro.Rewrite("$_.maxCpuSpeed")
	return maxCpuSpeed
}

// PaintRequestsPerSecond prop
// js:"paintRequestsPerSecond"
func (*PerfWidgetExternal) PaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	macro.Rewrite("$_.paintRequestsPerSecond")
	return paintRequestsPerSecond
}

// PerformanceCounter prop
// js:"performanceCounter"
func (*PerfWidgetExternal) PerformanceCounter() (performanceCounter uint64) {
	macro.Rewrite("$_.performanceCounter")
	return performanceCounter
}

// PerformanceCounterFrequency prop
// js:"performanceCounterFrequency"
func (*PerfWidgetExternal) PerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	macro.Rewrite("$_.performanceCounterFrequency")
	return performanceCounterFrequency
}
