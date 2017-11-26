package calc

//go:generate peg -inline structtag.peg

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	rulee
	rulee1
	rulee2
	rulee3
	rulee4
	rulevalue
	ruleadd
	ruleminus
	rulemultiply
	ruledivide
	rulemodulus
	ruleexponentiation
	ruleopen
	ruleclose
	rulesp
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	rulePegText
	ruleAction7
)

var rul3s = [...]string{
	"Unknown",
	"e",
	"e1",
	"e2",
	"e3",
	"e4",
	"value",
	"add",
	"minus",
	"multiply",
	"divide",
	"modulus",
	"exponentiation",
	"open",
	"close",
	"sp",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"PegText",
	"Action7",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type Calculator struct {
	Expression

	Buffer string
	buffer []rune
	rules  [25]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *Calculator) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *Calculator) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *Calculator
	max token32
}

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *Calculator) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *Calculator) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:
			p.AddOperator(TypeAdd)
		case ruleAction1:
			p.AddOperator(TypeSubtract)
		case ruleAction2:
			p.AddOperator(TypeMultiply)
		case ruleAction3:
			p.AddOperator(TypeDivide)
		case ruleAction4:
			p.AddOperator(TypeModulus)
		case ruleAction5:
			p.AddOperator(TypeExponentiation)
		case ruleAction6:
			p.AddOperator(TypeNegation)
		case ruleAction7:
			p.AddValue(buffer[begin:end])

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *Calculator) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 e <- <(sp e1 !.)> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				if !_rules[rulesp]() {
					goto l0
				}
				if !_rules[rulee1]() {
					goto l0
				}
				{
					position2, tokenIndex2 := position, tokenIndex
					if !matchDot() {
						goto l2
					}
					goto l0
				l2:
					position, tokenIndex = position2, tokenIndex2
				}
				add(rulee, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 e1 <- <(e2 ((add e2 Action0) / (minus e2 Action1))*)> */
		func() bool {
			position3, tokenIndex3 := position, tokenIndex
			{
				position4 := position
				if !_rules[rulee2]() {
					goto l3
				}
			l5:
				{
					position6, tokenIndex6 := position, tokenIndex
					{
						position7, tokenIndex7 := position, tokenIndex
						{
							position9 := position
							if buffer[position] != rune('+') {
								goto l8
							}
							position++
							if !_rules[rulesp]() {
								goto l8
							}
							add(ruleadd, position9)
						}
						if !_rules[rulee2]() {
							goto l8
						}
						{
							add(ruleAction0, position)
						}
						goto l7
					l8:
						position, tokenIndex = position7, tokenIndex7
						if !_rules[ruleminus]() {
							goto l6
						}
						if !_rules[rulee2]() {
							goto l6
						}
						{
							add(ruleAction1, position)
						}
					}
				l7:
					goto l5
				l6:
					position, tokenIndex = position6, tokenIndex6
				}
				add(rulee1, position4)
			}
			return true
		l3:
			position, tokenIndex = position3, tokenIndex3
			return false
		},
		/* 2 e2 <- <(e3 ((multiply e3 Action2) / (divide e3 Action3) / (modulus e3 Action4))*)> */
		func() bool {
			position12, tokenIndex12 := position, tokenIndex
			{
				position13 := position
				if !_rules[rulee3]() {
					goto l12
				}
			l14:
				{
					position15, tokenIndex15 := position, tokenIndex
					{
						position16, tokenIndex16 := position, tokenIndex
						{
							position18 := position
							if buffer[position] != rune('*') {
								goto l17
							}
							position++
							if !_rules[rulesp]() {
								goto l17
							}
							add(rulemultiply, position18)
						}
						if !_rules[rulee3]() {
							goto l17
						}
						{
							add(ruleAction2, position)
						}
						goto l16
					l17:
						position, tokenIndex = position16, tokenIndex16
						{
							position21 := position
							if buffer[position] != rune('/') {
								goto l20
							}
							position++
							if !_rules[rulesp]() {
								goto l20
							}
							add(ruledivide, position21)
						}
						if !_rules[rulee3]() {
							goto l20
						}
						{
							add(ruleAction3, position)
						}
						goto l16
					l20:
						position, tokenIndex = position16, tokenIndex16
						{
							position23 := position
							if buffer[position] != rune('%') {
								goto l15
							}
							position++
							if !_rules[rulesp]() {
								goto l15
							}
							add(rulemodulus, position23)
						}
						if !_rules[rulee3]() {
							goto l15
						}
						{
							add(ruleAction4, position)
						}
					}
				l16:
					goto l14
				l15:
					position, tokenIndex = position15, tokenIndex15
				}
				add(rulee2, position13)
			}
			return true
		l12:
			position, tokenIndex = position12, tokenIndex12
			return false
		},
		/* 3 e3 <- <(e4 (exponentiation e4 Action5)*)> */
		func() bool {
			position25, tokenIndex25 := position, tokenIndex
			{
				position26 := position
				if !_rules[rulee4]() {
					goto l25
				}
			l27:
				{
					position28, tokenIndex28 := position, tokenIndex
					{
						position29 := position
						if buffer[position] != rune('^') {
							goto l28
						}
						position++
						if !_rules[rulesp]() {
							goto l28
						}
						add(ruleexponentiation, position29)
					}
					if !_rules[rulee4]() {
						goto l28
					}
					{
						add(ruleAction5, position)
					}
					goto l27
				l28:
					position, tokenIndex = position28, tokenIndex28
				}
				add(rulee3, position26)
			}
			return true
		l25:
			position, tokenIndex = position25, tokenIndex25
			return false
		},
		/* 4 e4 <- <((minus value Action6) / value)> */
		func() bool {
			position31, tokenIndex31 := position, tokenIndex
			{
				position32 := position
				{
					position33, tokenIndex33 := position, tokenIndex
					if !_rules[ruleminus]() {
						goto l34
					}
					if !_rules[rulevalue]() {
						goto l34
					}
					{
						add(ruleAction6, position)
					}
					goto l33
				l34:
					position, tokenIndex = position33, tokenIndex33
					if !_rules[rulevalue]() {
						goto l31
					}
				}
			l33:
				add(rulee4, position32)
			}
			return true
		l31:
			position, tokenIndex = position31, tokenIndex31
			return false
		},
		/* 5 value <- <((<[0-9]+> sp Action7) / (open e1 close))> */
		func() bool {
			position36, tokenIndex36 := position, tokenIndex
			{
				position37 := position
				{
					position38, tokenIndex38 := position, tokenIndex
					{
						position40 := position
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l39
						}
						position++
					l41:
						{
							position42, tokenIndex42 := position, tokenIndex
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l42
							}
							position++
							goto l41
						l42:
							position, tokenIndex = position42, tokenIndex42
						}
						add(rulePegText, position40)
					}
					if !_rules[rulesp]() {
						goto l39
					}
					{
						add(ruleAction7, position)
					}
					goto l38
				l39:
					position, tokenIndex = position38, tokenIndex38
					{
						position44 := position
						if buffer[position] != rune('(') {
							goto l36
						}
						position++
						if !_rules[rulesp]() {
							goto l36
						}
						add(ruleopen, position44)
					}
					if !_rules[rulee1]() {
						goto l36
					}
					{
						position45 := position
						if buffer[position] != rune(')') {
							goto l36
						}
						position++
						if !_rules[rulesp]() {
							goto l36
						}
						add(ruleclose, position45)
					}
				}
			l38:
				add(rulevalue, position37)
			}
			return true
		l36:
			position, tokenIndex = position36, tokenIndex36
			return false
		},
		/* 6 add <- <('+' sp)> */
		nil,
		/* 7 minus <- <('-' sp)> */
		func() bool {
			position47, tokenIndex47 := position, tokenIndex
			{
				position48 := position
				if buffer[position] != rune('-') {
					goto l47
				}
				position++
				if !_rules[rulesp]() {
					goto l47
				}
				add(ruleminus, position48)
			}
			return true
		l47:
			position, tokenIndex = position47, tokenIndex47
			return false
		},
		/* 8 multiply <- <('*' sp)> */
		nil,
		/* 9 divide <- <('/' sp)> */
		nil,
		/* 10 modulus <- <('%' sp)> */
		nil,
		/* 11 exponentiation <- <('^' sp)> */
		nil,
		/* 12 open <- <('(' sp)> */
		nil,
		/* 13 close <- <(')' sp)> */
		nil,
		/* 14 sp <- <(' ' / '\t')*> */
		func() bool {
			{
				position56 := position
			l57:
				{
					position58, tokenIndex58 := position, tokenIndex
					{
						position59, tokenIndex59 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l60
						}
						position++
						goto l59
					l60:
						position, tokenIndex = position59, tokenIndex59
						if buffer[position] != rune('\t') {
							goto l58
						}
						position++
					}
				l59:
					goto l57
				l58:
					position, tokenIndex = position58, tokenIndex58
				}
				add(rulesp, position56)
			}
			return true
		},
		/* 16 Action0 <- <{ p.AddOperator(TypeAdd) }> */
		nil,
		/* 17 Action1 <- <{ p.AddOperator(TypeSubtract) }> */
		nil,
		/* 18 Action2 <- <{ p.AddOperator(TypeMultiply) }> */
		nil,
		/* 19 Action3 <- <{ p.AddOperator(TypeDivide) }> */
		nil,
		/* 20 Action4 <- <{ p.AddOperator(TypeModulus) }> */
		nil,
		/* 21 Action5 <- <{ p.AddOperator(TypeExponentiation) }> */
		nil,
		/* 22 Action6 <- <{ p.AddOperator(TypeNegation) }> */
		nil,
		nil,
		/* 24 Action7 <- <{ p.AddValue(buffer[begin:end]) }> */
		nil,
	}
	p.rules = _rules
}
