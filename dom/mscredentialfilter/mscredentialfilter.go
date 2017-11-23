package mscredentialfilter

import "github.com/matthewmueller/golly/dom/mscredentialspec"

type MSCredentialFilter struct {
	accept *[]*mscredentialspec.MSCredentialSpec
}
