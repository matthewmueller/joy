package msutilization

type MSUtilization struct {
	bandwidthEstimation       *uint64
	bandwidthEstimationAvg    *uint64
	bandwidthEstimationMax    *uint64
	bandwidthEstimationMin    *uint64
	bandwidthEstimationStdDev *uint64
	packets                   *uint64
}
