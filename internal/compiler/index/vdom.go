package index

import (
	"errors"

	"github.com/matthewmueller/golly/internal/compiler/def"
)

// SetVDOMPragma sets up JSX
func (i *Index) SetVDOMPragma(pragma string) {
	i.vdom.Pragma = pragma
}

// SetVDOMFile sets the vdom file
func (i *Index) SetVDOMFile(def def.Definition) {
	i.vdom.File = def
}

// VDOMFile gets the vdom file
func (i *Index) VDOMFile() (def.Definition, error) {
	if i.vdom.File == nil {
		return nil, errors.New("JSX not setup. Please use jsx.Use(pragma, filepath) in an init()")
	}
	return i.vdom.File, nil
}

// VDOMPragma gets our jsx settings if we've specified them
func (i *Index) VDOMPragma() (string, error) {
	if i.vdom.Pragma == "" {
		return "", errors.New("JSX not setup. Please use jsx.Use(pragma, filepath) in an init()")
	}
	return i.vdom.Pragma, nil
}
