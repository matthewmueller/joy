package jsast

// Parse javascript src
func Parse(src string) (IExpression, error) {
	return CreateCallExpression(
		CreateMemberExpression(
			CreateIdentifier("document"),
			CreateIdentifier("window"),
			false,
		),
		[]IExpression{},
	), nil
}
