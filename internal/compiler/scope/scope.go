package scope

import "go/ast"

// Scope type
type Scope struct {
	Owner ast.Node
	Outer *Scope
	scope *ast.Scope
}

// File scope is the outermost scope (I think... :-P)
// func File(scope *ast.Scope) *Scope {
// 	return &Scope{
// 		scope: scope,
// 	}
// }

// New Scope based on an old scope
func New(owner ast.Node) *Scope {
	return &Scope{
		Owner: owner,
		scope: ast.NewScope(nil),
	}
}

// Within looks up a name within a scope (doesn't traverse upward)
func (s *Scope) Within(name string) *ast.Object {
	return s.scope.Lookup(name)
}

// Lookup name in scope and all parent scopes
func (s *Scope) Lookup(name string) *ast.Object {
	o := s.scope.Lookup(name)
	if o != nil {
		return o
	}
	if s.Outer == nil {
		return nil
	}
	return s.Outer.Lookup(name)
}

// Insert attempts to insert a named object obj into the scope s.
// If the scope already contains an object alt with the same name,
// Insert leaves the scope unchanged and returns alt. Otherwise
// it inserts obj and returns nil.
func (s *Scope) Insert(obj *ast.Object) (alt *ast.Object) {
	return s.scope.Insert(obj)
}

// Debugging support
func (s *Scope) String() string {
	return s.scope.String()
}
