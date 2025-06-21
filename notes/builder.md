# Builder Pattern

- `Builder Pattern` is used to create a object piece by piece and not all at once with a big ass constructor.
- `Fluent` functions are functions which return the same interface back so that we can chain methods

```go
func (i *FluentInterface) CallFunc(name string) *FluentInterface {
 fmt.Println(name)
 return i
}

func main() {
 i := FluentInterface{};
 i.CallFunc("Hello").CallFunc("World")
}
```

- As you can see here, the two functions can be chained as they are passing the interface
- In the `HTMLBuilder` example, we were trying to create a struct through which we can construct an HTML component without much code calls. In the end, we were able to create a whole list using just this -

```go
fb.AddChildFluent("li", "item 1").AddChildFluent("li", "item 2")
```

## Builder Facets

- For creating a builder for a struct, we can also divide the work into multiple builder function and then aggregate them.

#### Compositon notes

- If there is a struct `Parent` and another pair of structs `Child1` and `Child2` and this is the setup -

```go
type Parent struct{}
type Child1 struct{
 Parent
}
type Child2 struct{
 Parent
}

func(p *Parent) ParentMethod() {}
func(c *Child1) ChildMethod1() {}
func(c *Child2) ChildMethod2() {}
```

- Objects of `Parent` will only have access to `ParentMethod()`
- Objects of `Child1` will have access to `ParentMethod()` and `ChildMethod1()`
- Objects of `Child2` will have access to `ParentMethod()` and `ChildMethod2()`

#### Person Builder Example

- In the `PersonBuilder` example, we have two sets of data - Address related data and Job related data.
- So we created three builders - `PersonBuilder`, `PersonJobBuilder` and `PersonAddressBuilder`.
- `PersonBuilder` had methods to convert to either of the two children to add data of that scope and a method to return the `Person`
- `PersonJobBuilder` and `PersonAddressBuilder` had functions to add in fields of their specific scope and they could not access the other fields until they switch to the other one using the `PersonBuilder`'s method.
- ![[Builder Pattern 2025-06-20 12.50.41.excalidraw]]
- This example shows that it's possible to have multiple builders for different aspects of an object and make them share a parent builder which can be used to switch between different builders.

## Builder Params

- Now, we know how to create builders but how do we utilise them. Builders can also be used to give an abstraction over the base class.
- We will discuss here about how to implement functions directly with Builders rather than the base object
- For example, there's an `Email` interface and there is a function `sendEmailImpl` which sends the mail using an object of Email.
- Here's how we can refactor the code to add Builder Pattern and a layer of abstraction

```go
type Email struct {}
func sendEmailImpl(email *Email) {}

type Builder struct {
 email Email
}
func (b* Builder) From(email string) {}
// similar functions to help create Email

type build func(*Builder)

func SendMail(action build) {
 emailBuilder := Builder{}
 action(&emailBuilder)
 sendMailEmailImpl(&emailBuilder.email)
}

func main() {
 SendMail(func(eb *Builder) {
  eb.
   From("abc@gmail.com").
   To("def@yahoo.com").
   Subject("First email").
   Body("Hello, how are you?")
 })
}
```

- Inside the `build` function, we can create the object using the builder function and as you can see in the `SendEmail` function, we get the built email after doing `action` and then we can share the object, all while making sure that the internal object is not touched by the `SendEmail` function
- Also this can be used to validate the fields of the object with the functions responsible for building

#### Error handling in Builder

- If the builder methods have validation then it's important to make sure to handle errors effectively
- This can be done by having an err field of type error.
- As soon as there's an error in running a builder method, we set the err field to that error.
- And in every method following that method we just return the builder without doing anything.

```go
type Builder struct {
 email Email
 err   error
}

func (b *Builder) From(from string) *Builder {
 if b.err != nil {
  return b
 }  //  This makes sure that if there was any error in the previous calls, we will not do any processing in this call
 if !strings.Contains(from, "@") {
  b.err = fmt.Errorf("From should have an @")
  // this sets the error and returns
  return b
 }
 b.email.From = from
 return b
}
```

- As you can see here, if there is any error, we will not do any further processing and just return the builder with the error.

## Functional Builder

- The idea behind a functional builder is the delayed application of the changes.
- Instead of doing the changes right away, we keep a list of the actions that need to be taken for building an object
- Once build is called, then only we create all the actions and create the object.
- I personally think this approach is better.
- This can also be used as builder params.

```go
type personAMod func(*PersonA)
type PersonABuilder struct {
 actions []personAMod
}

func (b *PersonABuilder) Called(name string) *PersonABuilder {
 b.actions = append(b.actions, func(pa *PersonA) {
  pa.name = name
 })

 return b
}

func (b *PersonABuilder) Build() *PersonA {
 p := PersonA{}
 for _, a := range b.actions {
  a(&p)
 }
 return &p
}

func Introduce(action personBuild) {
 pb := PersonABuilder{}
 action(&pb)
 // at this point, all the actions required to create a function have been added
 introducePerson(pb.Build())
}

Introduce(func(pa *PersonABuilder) {
 pa.Called("Shreyash").Is("Developer")
})

```
