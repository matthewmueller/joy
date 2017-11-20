package rtcsrtpsdesparameters

import "github.com/matthewmueller/golly/internal/gendom/dom/rtcsrtpkeyparam"

type RTCSrtpSdesParameters struct {
	cryptoSuite   *string
	keyParams     *[]*rtcsrtpkeyparam.RTCSrtpKeyParam
	sessionParams *[]string
	tag           *uint8
}
