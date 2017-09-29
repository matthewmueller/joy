package golang

import "go/ast"

func variablePairs(n ast.Node) {
	switch t := n.(type) {
	case *ast.ValueSpec:
		valueSpecPairs(t)
	case *ast.AssignStmt:
		assignStmtPairs(t)
	default:
		unhandled("variablePairs", t)
	}
}

func valueSpecPairs(n *ast.ValueSpec) {

}

func assignStmtPairs(n *ast.AssignStmt) {

}
