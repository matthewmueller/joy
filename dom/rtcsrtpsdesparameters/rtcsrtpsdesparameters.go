package rtcsrtpsdesparameters

import "github.com/matthewmueller/joy/dom/rtcsrtpkeyparam"

type RTCSrtpSdesParameters struct {
	cryptoSuite   *string
	keyParams     *[]*rtcsrtpkeyparam.RTCSrtpKeyParam
	sessionParams *[]string
	tag           *uint8
}
