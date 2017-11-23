package rtcsrtpsdesparameters

import "github.com/matthewmueller/golly/dom2/rtcsrtpkeyparam"

type RTCSrtpSdesParameters struct {
	cryptoSuite   *string
	keyParams     *[]*rtcsrtpkeyparam.RTCSrtpKeyParam
	sessionParams *[]string
	tag           *uint8
}
