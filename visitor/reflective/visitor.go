package reflective

import (
	"fmt"
	"strings"
)

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(e Expression, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionExpression); ok {
		sb.WriteRune('(')
		Print(ae.left, sb)
		sb.WriteRune('+')
		Print(ae.right, sb)
		sb.WriteRune(')')
	}
}

func TestReflectiveVisitor() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		left: &DoubleExpression{value: 1},
		right: &AdditionExpression{
			left:  &DoubleExpression{value: 2},
			right: &DoubleExpression{value: 3},
		},
	}
	sb := new(strings.Builder)
	Print(e, sb)
	fmt.Println(sb.String())
}
