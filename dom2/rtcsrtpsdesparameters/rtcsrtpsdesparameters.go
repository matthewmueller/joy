package rtcsrtpsdesparameters

type RTCSrtpSdesParameters struct {
	cryptoSuite   *string
	keyParams     *[]*rtcsrtpkeyparam.RTCSrtpKeyParam
	sessionParams *[]string
	tag           *uint8
}
