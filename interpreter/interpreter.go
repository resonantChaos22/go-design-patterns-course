package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Lexing Process
type TokenType int

const (
	Int TokenType = iota
	Plus
	Minus
	LParen
	RParen
)

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
	var result []Token
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{Type: Plus, Text: "+"})
			continue
		case '-':
			result = append(result, Token{Type: Minus, Text: "-"})
			continue
		case '(':
			result = append(result, Token{Type: LParen, Text: "("})
			continue
		case ')':
			result = append(result, Token{Type: RParen, Text: ")"})
			continue

		default:
			if unicode.IsDigit(rune(input[i])) {
				sb := strings.Builder{}
				for j := i; j < len(input); j++ {
					if unicode.IsDigit(rune(input[j])) {
						sb.WriteRune(rune(input[j]))
						i++
					} else {
						result = append(result, Token{Type: Int, Text: sb.String()})
						i--
						break
					}
				}
				continue
			}
		}
	}

	return result
}

// Parsing code
type Element interface {
	Value() int
}

type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

func NewInteger(value int) *Integer {
	return &Integer{
		value: value,
	}
}

type Operation int

const (
	Addition Operation = iota
	Subtraction
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func (b *BinaryOperation) Value() int {
	switch b.Type {
	case Addition:
		return b.Left.Value() + b.Right.Value()
	case Subtraction:
		return b.Left.Value() - b.Right.Value()

	default:
		panic("Unsupported Operation")
	}
}

func Parse(tokens []Token) Element {
	result := BinaryOperation{}
	haveLhs := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case Int:
			n, _ := strconv.Atoi(token.Text)
			integer := NewInteger(n)
			if !haveLhs {
				result.Left = integer
				haveLhs = true
			} else {
				result.Right = integer
			}
			continue

		case Plus:
			result.Type = Addition
			continue

		case Minus:
			result.Type = Subtraction
			continue

		case LParen:
			j := i
			numParantheses := -1
			for ; j < len(tokens); j++ {
				if tokens[j].Type == RParen {
					if numParantheses == 0 {
						break
					} else {
						numParantheses--
					}
				}
				if tokens[j].Type == LParen {
					numParantheses++
				}
			}
			var subexp []Token
			for k := i + 1; k < j; k++ {
				subexp = append(subexp, tokens[k])
			}
			element := Parse(subexp)
			if !haveLhs {
				result.Left = element
				haveLhs = true
			} else {
				result.Right = element
			}
			i = j
		}
	}

	return &result
}

func TestInterpreter() {
	input := "((13+4)-(12+1)) + (12 + 7)"
	tokens := Lex(input)
	for _, token := range tokens {
		fmt.Printf("%s", token.Text)
	}
	fmt.Println()

	parsed := Parse(tokens)
	fmt.Printf("%s = %d\n", input, parsed.Value())
}
