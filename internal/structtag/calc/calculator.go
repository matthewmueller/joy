package calc

import "math/big"

type Type uint8

const (
	TypeNumber Type = iota
	TypeNegation
	TypeAdd
	TypeSubtract
	TypeMultiply
	TypeDivide
	TypeModulus
	TypeExponentiation
)

type ByteCode struct {
	T     Type
	Value *big.Int
}

func (code *ByteCode) String() string {
	switch code.T {
	case TypeNumber:
		return code.Value.String()
	case TypeAdd:
		return "+"
	case TypeNegation, TypeSubtract:
		return "-"
	case TypeMultiply:
		return "*"
	case TypeDivide:
		return "/"
	case TypeModulus:
		return "%"
	case TypeExponentiation:
		return "^"
	}
	return ""
}

type Expression struct {
	Code []ByteCode
	Top  int
}

func (e *Expression) Init(expression string) {
	e.Code = make([]ByteCode, len(expression))
}

func (e *Expression) AddOperator(operator Type) {
	code, top := e.Code, e.Top
	e.Top++
	code[top].T = operator
}

func (e *Expression) AddValue(value string) {
	code, top := e.Code, e.Top
	e.Top++
	code[top].Value = new(big.Int)
	code[top].Value.SetString(value, 10)
}

func (e *Expression) Evaluate() *big.Int {
	stack, top := make([]big.Int, len(e.Code)), 0
	for _, code := range e.Code[0:e.Top] {
		switch code.T {
		case TypeNumber:
			stack[top].Set(code.Value)
			top++
			continue
		case TypeNegation:
			a := &stack[top-1]
			a.Neg(a)
			continue
		}
		a, b := &stack[top-2], &stack[top-1]
		top--
		switch code.T {
		case TypeAdd:
			a.Add(a, b)
		case TypeSubtract:
			a.Sub(a, b)
		case TypeMultiply:
			a.Mul(a, b)
		case TypeDivide:
			a.Div(a, b)
		case TypeModulus:
			a.Mod(a, b)
		case TypeExponentiation:
			a.Exp(a, b, nil)
		}
	}
	return &stack[0]
}
