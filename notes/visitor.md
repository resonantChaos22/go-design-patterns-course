- This is an alternative to `Iterator`
- `Visitor Pattern` is a pattern where a component (visitor) is allowed to traverse the entire hierarchy of types. Implemented by propagating a single `Accept()` method throughout the entire hierarchy.

## Intrusive Visitor

- Intrusive Visitor is the case when we break OCP and add methods for the structs being "visited" for them being visited.

```go
type Expression interface {
 //  didnt originally have Print()
 Print(sb *strings.Builder)
}

type DoubleExpression struct {
 value float64
}
//  didnt originally have Print()
func (d *DoubleExpression) Print(sb *strings.Builder) {
 sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
 left, right Expression
}

//  didnt originally have Print()
func (a *AdditionExpression) Print(sb *strings.Builder) {
 sb.WriteRune('(')
 a.left.Print(sb)
 sb.WriteRune('+')
 a.right.Print(sb)
 sb.WriteRune(')')
}
```

- In this case `sb` is the visitor as it's visiting all the structs in the hierarchy

## Reflective Visitor

- In `Reflective Visitor`, we try to typecast the interface to a struct and then perform actions accordingly. This makes it so that the OCP is not breaking at the moment as we are not changing the existing implementation.

```go
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

```

- The issue we run into is, what if there is a third struct that is implemented as an Expression? Then, we will have to add another condition here and there is a possibility that we might forget to add it.
- Also, it causes an issue if we want to implement another visitor, something like get solution

## Dispatch

- It's a concept that is meant for choosing the function to call
- For `foo.Bar()`, we call a `Single Dispatch` where it depends on the name of request (`Bar()`) and the type of receiver (`foo`)
- There is a concept called `Double Dispatch` which depends on the name of request and type of two receivers ( type of visitor and type of element being visited ). This is something like function overload but go does not implement it.

## Classic Visitor

- The classic visitor pattern is based on double dispatch where the recursive function is called with two receivers, the structure that is being visited and the visitor that is being visited.
- Since, golang does not support function overloading, we have to do a small workaround. Doing this workaround would mean to break the OCP but this will add the functionality of `Visitor Pattern`once and for all so that for any other visitors, the existing interface wont change and we wont forget the implementation of a particular struct as well

```go
type ExpressionVisitor interface {
 VisitDoubleExpression(e *DoubleExpression)
 VisitAddtionExpression(e *AdditionExpression)
}

type Expression interface {
 Accept(ev ExpressionVisitor)
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
 ev.VisitDoubleExpression(d)
}
```

- So this `Accept()` method is implemented for both the structs and what it does is to call the struct-specific expression visitor as you can see here with `VisitDoubleExpression` with the current struct.

```go
func (ep *ExpressionPrinter) VisitAddtionExpression(e *AdditionExpression) {
 ep.sb.WriteRune('(')
 e.left.Accept(ep)
 ep.sb.WriteRune('+')
 e.right.Accept(ep)
 ep.sb.WriteRune(')')
}
```

- For example for `ExpressionPrinter`, we would implement it like this. Now we know that both `left` and `right` implement `Expression` so they will have `Accept` which will call the struct specific function to get the actual value of that and everything will work out recursively.
- Take example of this -

```go
e := &AdditionExpression{
 left: &DoubleExpression{value: 1},
 right: &AdditionExpression{
  left:  &DoubleExpression{value: 2},
  right: &DoubleExpression{value: 3},
 },
}
e.Accept(NewExpressionPrinter())
```

- On calling the accept these things happen
- `VisitAdditionExpression` is called
 	- Goes to `e.left.Accept` which calls `VisitDoubleExpression` on the {value: 1} and thus the string will be `(1+`
 	- Goes to `e.right.Accept` which calls `VisitAdditionExpression` on the left:2, right: 3.
  		- Goes to `e.left.Accept` which calls `VisitDoubleExpression` on the {value: 2} and the string will be `(1+(2+`
  		- Goes to `e.right.Accept` which calls `VisitDoubleExpression` on the {value: 3} and the string will be `(1+(2+3)`
 	- Then `)` is added and the string is returned - `(1+(2+3))`
- To add another `Visitor` for evaluating the expression, we can do that -

```go
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
```

- At every visit, it gets the `Value` of that expression.
