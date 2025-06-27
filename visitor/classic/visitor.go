package classic

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression)
	VisitAddtionExpression(e *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAddtionExpression(a)
}

type ExpressionPrinter struct {
	sb *strings.Builder
}

func (ep *ExpressionPrinter) VisitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.value))
}
func (ep *ExpressionPrinter) VisitAddtionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.left.Accept(ep)
	ep.sb.WriteRune('+')
	e.right.Accept(ep)
	ep.sb.WriteRune(')')
}
func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{
		sb: new(strings.Builder),
	}
}

// TODO: Implement Expression Evaluator
type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(e *DoubleExpression) {
	ee.result = e.value
}
func (ee *ExpressionEvaluator) VisitAddtionExpression(e *AdditionExpression) {
	e.left.Accept(ee)
	x := ee.result
	e.right.Accept(ee)
	x += ee.result
	ee.result = x
}
func (ee *ExpressionEvaluator) Value() float64 {
	return ee.result
}

func NewExpressionEvaluator() *ExpressionEvaluator {
	return &ExpressionEvaluator{}
}

func TestClassicVisitor() {
	// 1 + (2 + 3)
	e := &AdditionExpression{
		left: &DoubleExpression{value: 1},
		right: &AdditionExpression{
			left:  &DoubleExpression{value: 2},
			right: &DoubleExpression{value: 3},
		},
	}

	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := NewExpressionEvaluator()
	e.Accept(ee)

	fmt.Printf("%s = %g\n", ep, ee.Value())
}
