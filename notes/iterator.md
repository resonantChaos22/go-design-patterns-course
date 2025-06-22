- Iterator is a type that facilitates the traversal of a data structure
- Not idiomatic in go but still have usecases.
- Moves along the iterated collection and indicates when it has reached the end

## Example: iterating over the name fields in a Person

- There are 3 approaches to iterate over the `FirstName`, `MiddleName` and `LastName` fields of a person
- First- just create a function to return the 3 names as an array of strings.

```go
func (p *Person) Names() [3]string {
 return [3]string{p.FirstName, p.MiddleName, p.LastName}
}
```

- The issue with this is that this does not deal with the possibility that the middle name could be empty and while we can solve this by having more logic and making it a slice instead of an array, there are better ways.
- Second Way - Generator ie use channels and go routines to create a generator which sends out the names if they are valid.

```go
func (p *Person) NamesGenerator() <-chan string {
 out := make(chan string)
 go func() {
  defer close(out)
  out <- p.FirstName
  if len(p.MiddleName) > 0 {
   out <- p.MiddleName
  }
  out <- p.LastName
 }()

 return out
}
```

- The goroutine handles all the logic of which names to send and we can iterate over the channel that has been given as an output
- Third approach is not idiomatic to go but we can create a struct which functions just like a cpp iterator.

```go
type PersonNameIterator struct {
 person  *Person
 current int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
 return &PersonNameIterator{
  person:  person,
  current: -1,
 }
}
func (p *PersonNameIterator) MoveNext() bool {
 p.current++
 return p.current < 3
}
func (p *PersonNameIterator) Value() string {
 switch p.current {
 case 0:
  return p.person.FirstName
 case 1:
  return p.person.MiddleName
 case 2:
  return p.person.LastName

 default:
  panic("Not supported")
 }
}
for it := NewPersonNameIterator(&p); it.MoveNext(); {
 fmt.Printf("%s ", it.Value())
}
```

- It moves next until it can, then stops, the Value is given based on the position and additional logic could easily be added here to output data after validation
- Typically in Iterator Design Pattern, we talk about this approach only where we create a struct.

## Tree Traversal

- In tree traversal, we need to have a iterator struct if we want to iterate through the tree.
- We built `Inorder` Iterator which does exactly that -

```go
type Node struct {
 Value               int
 left, right, parent *Node
}

type InorderIterator struct {
 Current       *Node
 root          *Node
 returnedStart bool
}

func (i *InorderIterator) MoveNext() bool {
 if i.Current == nil {
  return false
 }
 if !i.returnedStart {
  i.returnedStart = true
  return true
 }

 if i.Current.right != nil {
  i.Current = i.Current.right
  for i.Current.left != nil {
   i.Current = i.Current.left
  }
  return true
 } else {
  p := i.Current.parent
  for p != nil && i.Current == p.right {
   i.Current = p
   p = p.parent
  }
  i.Current = p
  return i.Current != nil
 }
}

it := NewInorderIterator(root)
for it.MoveNext() {
 fmt.Printf("%d, ", it.Current.Value)
}
```

- During the initialization of the iterator, we make sure that the current is leftest node int the tree.
- Then the iterator just iterates through the tree like an inorder traversal and we can get the value at every point of the way.
- `returnedStart` is used to make sure that if we are visiting a node for the first time, we print it out first.
