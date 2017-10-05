package binding

import (
	"go/ast"
	"time"

	"github.com/apex/log"
)

//  0  *ast.ExprStmt {
//  1  .  X: *ast.CallExpr {
//  2  .  .  Fun: *ast.SelectorExpr {
//  3  .  .  .  X: *ast.CallExpr {
//  4  .  .  .  .  Fun: *ast.SelectorExpr {
//  5  .  .  .  .  .  X: *ast.CallExpr {
//  6  .  .  .  .  .  .  Fun: *ast.SelectorExpr {
//  7  .  .  .  .  .  .  .  X: *ast.Ident {
//  8  .  .  .  .  .  .  .  .  NamePos: 265
//  9  .  .  .  .  .  .  .  .  Name: "js"
//  10  .  .  .  .  .  .  .  }
//  11  .  .  .  .  .  .  .  Sel: *ast.Ident {
//  12  .  .  .  .  .  .  .  .  NamePos: 268
//  13  .  .  .  .  .  .  .  .  Name: "Promise"
//  14  .  .  .  .  .  .  .  }
//  15  .  .  .  .  .  .  }
//  16  .  .  .  .  .  .  Lparen: 275
//  17  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
//  18  .  .  .  .  .  .  .  0: *ast.BasicLit {
//  19  .  .  .  .  .  .  .  .  ValuePos: 276
//  20  .  .  .  .  .  .  .  .  Kind: STRING
//  21  .  .  .  .  .  .  .  .  Value: "\"window.fetch(url, options)\""
//  22  .  .  .  .  .  .  .  }
//  23  .  .  .  .  .  .  }
//  24  .  .  .  .  .  .  Ellipsis: 0
//  26  .  .  .  .  .  }
//  27  .  .  .  .  .  Sel: *ast.Ident {
//  28  .  .  .  .  .  .  NamePos: 306
//  29  .  .  .  .  .  .  Name: "Then"
//  30  .  .  .  .  .  }
//  31  .  .  .  .  }
//  32  .  .  .  .  Lparen: 310
//  33  .  .  .  .  Args: []ast.Expr (len = 1) {
//  34  .  .  .  .  .  0: *ast.Ident {
//  35  .  .  .  .  .  .  NamePos: 311
//  36  .  .  .  .  .  .  Name: "res"
// 161  .  .  .  .  .  }
// 162  .  .  .  .  }
// 163  .  .  .  .  Ellipsis: 0
// 165  .  .  .  }
// 166  .  .  .  Sel: *ast.Ident {
// 167  .  .  .  .  NamePos: 316
// 168  .  .  .  .  Name: "Catch"
// 169  .  .  .  }
// 170  .  .  }
// 171  .  .  Lparen: 321
// 172  .  .  Args: []ast.Expr (len = 1) {
// 173  .  .  .  0: *ast.Ident {
// 174  .  .  .  .  NamePos: 322
// 175  .  .  .  .  Name: "err"
// 193  .  .  .  }
// 194  .  .  }
// 195  .  .  Ellipsis: 0
// 197  .  }
// 198  }

// Promise fn
func Promise(n ast.Stmt) ([]interface{}, error) {
	start := time.Now()
	xs, ok := n.(*ast.ExprStmt)
	if !ok {
		return nil, nil
	}

	cx, ok := xs.X.(*ast.CallExpr)
	if !ok {
		return nil, nil
	}

	s, ok := cx.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}

	cx, ok = s.X.(*ast.CallExpr)
	if !ok {
		return nil, nil
	}

	s, ok = cx.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}

	cx, ok = s.X.(*ast.CallExpr)
	if !ok {
		return nil, nil
	}

	s, ok = cx.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}

	id, ok := s.X.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	if id.Name != "js" || s.Sel.Name != "Promise" {
		return nil, nil
	}

	log.Infof("got promise! %s", time.Since(start))
	return nil, nil
}
