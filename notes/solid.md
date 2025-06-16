
## Single Responsibility Principle (SRP)

- A class should have one primary responsibility and as a result should have one reason to change that is related to the primary responsibilities.
- `Separation Of Concerns` - Make sure that one package does one job only. For example, if you have a Journal Class, it should only be used to handle the journal.
- In case we want to add persistence (ie save journals to a file), you should create a separate package which can be used to save all kinds of objects.
- So later in the project, when you want to persist something again, you can use that package which would make further development better

## Open-Closed Principle (OCP)

- A class should be open for extension but closed for modification.
- `Specification Pattern` - Enterprise

```go
type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
 result := make([]*Product, 0)
 for i, v := range products {
  if v.color == color {
   result = append(result, &products[i])
  }
 }

 return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
 result := make([]*Product, 0)
 for i, v := range products {
  if v.size == size {
   result = append(result, &products[i])
  }
 }

 return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
 result := make([]*Product, 0)
 for i, v := range products {
  if v.size == size && v.color == color {
   result = append(result, &products[i])
  }
 }

 return result
}
```

- Take this code for example, if we only define the `FilterByColor`, to further add other filters, we would need to write a modification of the previously written code.
- This would lead to issues where one modification can lead to bugs. To prevent that, it's always better to write code such that there is no modification, only extension of the code.
- A better way to write this -

```go
type Specification interface {
 IsSatisfied(*Product) bool
}

type ColorSpecification struct {
 color Color
}

func (spec *ColorSpecification) IsSatisfied(product *Product) bool {
 return product.color == spec.color
}

type SizeSpecification struct {
 size Size
}

func (spec *SizeSpecification) IsSatisfied(product *Product) bool {
 return product.size == spec.size
}

type AndSpecification struct {
 specA Specification
 specB Specification
}

func (spec *AndSpecification) IsSatisfied(product *Product) bool {
 return spec.specA.IsSatisfied(product) && spec.specB.IsSatisfied(product)
}

type BetterFilter struct{}

func (bf BetterFilter) FilterValue(products []Product, spec Specification) []*Product {
 result := make([]*Product, 0)
 for i, v := range products {
  if spec.IsSatisfied(&v) {
   result = append(result, &products[i])
  }
 }

 return result
}

```

- This is also known as the `Specification Pattern` where we filter based on a `IsSatisfied` function.
- In conjunction to `AndSpecification`, this pattern can be used to create as many specifications/ filters as required without ever modifying the filter logic

## Liskov Substitution Principle

- If there's an API that works well with a Base Class, it should work well with the derived class as well
- Not applicable to go as there are no derived classes.
- For the case of golang, if we are using an interface, we have to make sure that the methods that we are making do not break any function that is using that interface.

## Interface Segragation Principle

- Dont put everything into a single interface, instead break it into smaller interfaces.
- Example -

```go
type Machine interface {
 Print(Document)
 Fax(Document)
 Scan(Document)
}

```

- This interface could be implemented by a `MultiFunctionPrinter` but it can not be implemented by `OldPrinter` which only supports `Print`
- If we try to forcefully define other methods for it, it could lead to confusion and bugs in the future.
- Instead do this -

```go

type Printer interface {
 Print(Document)
}

type Faxer interface {
 Fax(Document)
}

type Scanner interface {
 Scan(Document)
}

type MultiFunctionDevice interface {
 Printer
 Faxer
 Scanner
}
```

- This will provide to be a more robust solution

## Dependency Inversion Principle

- High level modules should not depend on Low Level Modules, they should both depend on abstractions
- Example -

```go
// low level module (basically data)
type Relationships struct {
 relations []Info
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
 r.relations = append(r.relations, Info{parent, Parent, child})
 r.relations = append(r.relations, Info{child, Child, parent})
}

// high level module (functions on data)
type Research struct {
 relationships Relationships
}

func (r *Research) Investigate() {
 relations := r.relationships.relations
 for _, rel := range relations {
  if rel.relationship == Parent && rel.from.name == "John" {
   fmt.Printf("%s has a child called %s\n", rel.from.name, rel.to.name)
  }
 }
}
```

- What happens if the implementation of low-level module changes, ie we change from in-memory to db or something else
- Then `Investigate()` function breaks as we are directly referencing the fields of the low level module
- To fix this, we will have to make changes in Research as well as the low level module which would lead to further work down the line.
- So, a better approach is this -

```go
type RelationshipBrowser interface {
 FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
 children := make([]*Person, 0)

 for _, rel := range r.relations {
  if rel.relationship == Parent && rel.from.name == name {
   children = append(children, rel.to)
  }
 }

 return children
}

type Research struct {
 browser RelationshipBrowser
}

func (r *Research) Investigate() {
 for _, child := range r.browser.FindAllChildrenOf("John") {
  fmt.Printf("Johb has a child called %s\n", child.name)
 }
}
```

- In this case, in `Investigate` function, we are no longer dependent on how the relationships are stored
- Whenever, the low-level implementation is updated, we would know that we only need to modify the methods of the low-level implementation
