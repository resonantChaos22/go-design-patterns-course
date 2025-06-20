## Factory Function

- A function that returns the instance of the struct that you want to create.
- This is useful due to many reasons like -
  - If there are default values in the struct
  - If there is some processing that needs to be done with some values every time that struct is to be created.
  - If there needs to be some validation

## Interface Factory

- The factory function returns a interface instead of the struct to encapsulate the information and abstract it so that when using the object from the Factory Function anywhere we just use the functions available to us rather than modifying the underlying struct by mistake.
- Moreover, we can also return different type of structs which have different implementations of that function but do the same thing. For example, a `Shape` interface can take

## Factory Generator

- `Factory Generator` is used when there are different types of the base struct. For example - Employee

```go
type Employee struct {
 Name         string
 Postion      string
 AnnualIncome int
}
```

- In this case, there could be different type of employees like "Developer", "Manager", "QA Engineer", etc which will have a set annual income and position.
- So we can generate new employees using just the name when using the factory.
- There are two approaches to this - Functional Approach and Structural Approach.
- In the `Functional Approach`, we create a function which returns another function. Something like this -

```go
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
 return func(name string) *Employee {
  return &Employee{
   Name:         name,
   Postion:      position,
   AnnualIncome: annualIncome,
  }
 }
}

developerFactory := NewEmployeeFactory("Developer", 60000)
dev := developerFactory("Shreyash")
```

- And that's how we can have different factories for different employees and these factories will only require the name of the employee to create the employee.
- Another approach is the `Structural Approach` where we create a `EmployeeFactory` struct with the required information and then define a method `Create` which can help us get the new employee -

```go
type EmployeeFactory struct {
 Position     string
 AnnualIncome int
}

func NewEmployeeFactoryB(position string, annualIncome int) *EmployeeFactory {
 return &EmployeeFactory{
  Position:     position,
  AnnualIncome: annualIncome,
 }
}

func (f *EmployeeFactory) Create(name string) *Employee {
 return &Employee{Name: name, Postion: f.Position, AnnualIncome: f.AnnualIncome}
}

developerFactory := NewEmployeeFactoryB("Developer", 60000)
dev := developerFactory.Create("Shreyash")
```

- Here also, the same theory is applicable.
- Structural approach if we need a factory whose fields can be altered, for example, if we want to update the annualIncome of a developer.
- Functional Approach is better if we dont want to update any info whenever as it the function can be directly called and the developer who is using the factory does not need to be aware of the functions associated with the factory to prevent confusion.

## Prototype Factory

- A subtype of `Factory Generator` where instead of generating different employees with a factory, we generate the employees with preconfigured fields directly based on role

```go
type role int

const (
 Developer role = iota
 Manager
)

func NewEmployee(role role) *Employee {
 switch role {
 case Developer:
  return &Employee{Postion: "Developer", AnnualIncome: 60000}
 case Manager:
  return &Employee{Postion: "Manager", AnnualIncome: 80000}
 default:
  panic("No such role")
 }
}

m := NewEmployee(Manager)
m.Name = "Sam"
```
