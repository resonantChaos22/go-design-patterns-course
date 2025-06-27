package main

import (
	"fmt"
	"strings"
)

//	intrusive approach - overlooking the OCP, as we are adding a new method for printing

type Expression interface {
	Print(sb *strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.left.Print(sb)
	sb.WriteRune('+')
	a.right.Print(sb)
	sb.WriteRune(')')
}

func TestIntrusiveVisitor() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		left: &DoubleExpression{value: 1},
		right: &AdditionExpression{
			left:  &DoubleExpression{value: 2},
			right: &DoubleExpression{value: 3},
		},
	}

	sb := new(strings.Builder)
	e.Print(sb)
	fmt.Println(sb.String())
}
