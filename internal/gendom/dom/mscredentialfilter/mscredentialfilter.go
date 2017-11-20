package mscredentialfilter

import "github.com/matthewmueller/golly/internal/gendom/dom/mscredentialspec"

type MSCredentialFilter struct {
	accept *[]*mscredentialspec.MSCredentialSpec
}
