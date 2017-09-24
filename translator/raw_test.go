package translator_test

// func Golly(t *testing.T) {
// 	translator.Golly(
// 		&ast.CallExpr{
// 			Fun: &ast.SelectorExpr{
// 				X: &ast.Ident{
// 					Name: "js",
// 				},
// 				Sel: &ast.Ident{
// 					Name: "CreateMemberExpression",
// 				},
// 			},
// 			Args: []ast.Expr{
// 				&ast.CallExpr{
// 					Fun: &ast.SelectorExpr{
// 						X: &ast.Ident{
// 							Name: "js",
// 						},
// 						Sel: &ast.Ident{
// 							Name: "CreateIdentifier",
// 						},
// 					},
// 					Args: []ast.Expr{
// 						&ast.BasicLit{
// 							Kind:  token.STRING,
// 							Value: "\"window\"",
// 						},
// 					},
// 				},
// 				&ast.CallExpr{
// 					Fun: &ast.SelectorExpr{
// 						X: &ast.Ident{
// 							Name: "js",
// 						},
// 						Sel: &ast.Ident{
// 							Name: "CreateIdentifier",
// 						},
// 					},
// 					Args: []ast.Expr{
// 						&ast.BasicLit{
// 							Kind:  token.STRING,
// 							Value: "\"document\"",
// 						},
// 					},
// 				},
// 				&ast.Ident{
// 					Name: "false",
// 				},
// 			},
// 		},
// 	)
// }

// func Go(t *testing.T) {

// }
