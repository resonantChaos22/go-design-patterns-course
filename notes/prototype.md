`Prototype` is a partial or complete structure that can be copied or customized to be used

## Deep Copying

- The process of copying everything from one object to another.
- Problem - the pointers are copied as well, so both the objects share the same object of the pointer in memory.
- So, when modifying that object, remember to create new object and assign its pointer to the copied object.

## Copying through a function

- WE can also add a method to deep copy for the structs and that would work as well.

```go
type Address struct {
 StreetAddress string
 City          string
 Country       string
}

func (a *Address) DeepCopy() *Address {
 return &Address{
  StreetAddress: a.StreetAddress,
  City:          a.City,
  Country:       a.Country,
 }
}

type Person struct {
 Name    string
 Address *Address
 Friends []string
}

func (p *Person) DeepCopy() *Person {
 q := *p
 q.Address = p.Address.DeepCopy()
 copy(q.Friends, p.Friends)

 return &q
}
jane := john.DeepCopy()
```

- Here, Jane will be a `Deep Copy` of John and we can change any parameters of Jane without it affecting John.
- Also, do note the function `copy`. The function is used to copy one slice to another slice.
- This is tedious as we have had to do a lot of hard work

## Deep Copy through serialization

- When we serialize a struct into bytes, the serializer goes into the memory and copies everything (basically it copies the data that the pointer is pointing to as well). Then when we deserialize it into bytes, it creates a new object from those bytes which essentially is a `Deep Copy` of the previous object.

```go
func (p *Person) DeepCopy() *Person {
 b := bytes.Buffer{}
 e := gob.NewEncoder(&b)
 _ = e.Encode(p)

 d := gob.NewDecoder(&b)
 result := Person{}
 _ = d.Decode(&result)
 return &result
}
```

- Here, the person is a reconstructed Deep Copy of the person being passed to the method.
- This is the most efficient way to do `Deep Copy

## Prototype Factory

- We discussed Prototype Factory in the factory section but that was a very novice approach as we had to manually update the name after we got the prototype.
- What we can instead do is create prototypes (which will be partially filled objects) and then based on the requirement, we deep copy the prototype and add the required new information from a separate function. Example -

```go
var mainOfficeProto = Employee{
 Office: OfficeAddress{
  StreetAddress: "123 East Drive",
  City:          "London",
 },
}

func newEmployeeFromProto(proto *Employee, name string, suite int) *Employee {
 newEmployee := proto.DeepCopy()
 newEmployee.Name = name
 newEmployee.Office.Suite = suite

 return newEmployee
}

func NewEmployee(office office, name string, suite int) *Employee {
 switch office {
 case MainOffice:
  return newEmployeeFromProto(&mainOfficeProto, name, suite)

 case AuxOffice:
  return newEmployeeFromProto(&auxOfficeProto, name, suite)

 default:
  panic("No such office")
 }
}

john := NewEmployee(MainOffice, "John", 102)
```

- As we can see here, a new Employee is created as a deep copy from a proto and then the required information is updated and the new employee is returned
