package mscredentialfilter

import "github.com/matthewmueller/joy/dom/mscredentialspec"

type MSCredentialFilter struct {
	accept *[]*mscredentialspec.MSCredentialSpec
}
