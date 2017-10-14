package indexer

import "golang.org/x/tools/go/loader"
import "github.com/matthewmueller/golly/compiler/decl"

// Index struct
type Index struct{}

// New index
func New(program *loader.Program) (*Index, error) {
	return nil, nil
}

// Mains gets the main entrypoints of the program
func (i *Index) Mains() (decls []decl.Declaration, err error) {
	return decls, nil
}
