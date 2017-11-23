package mscredentialfilter

import "github.com/matthewmueller/golly/dom2/mscredentialspec"

type MSCredentialFilter struct {
	accept *[]*mscredentialspec.MSCredentialSpec
}
