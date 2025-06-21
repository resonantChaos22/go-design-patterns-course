- `Singleton Pattern` suggests that an object which takes a lot of time to load or there are other parameters which make it heavy should be instantiated only once and then a pointer should be used to refer to that.
- This can be done using go's `sync.Once` package

```go
func GetSingletonDB() *internalDatabase {
 once.Do(func() {
  fmt.Println("Initializing Database")
  db := internalDatabase{
   capitals: map[string]int{
    "Delhi":    12435323,
    "Seoul":    35432343,
    "New York": 53564454,
   },
  }
  internalDB = &db
 })
 return internalDB
}
fmt.Println(GetSingletonDB().GetCapital("Delhi"))
fmt.Println(GetSingletonDB().GetCapital("Seoul"))
```

- Now, even if the `GetSingletonDB` is called twice, it will only run once throughout the program's execution
- We could also have done this using `init()` method of the package but we also want to implement `laziness` ie the DB wont be initialized until we actually need it, so for that, we can only using sync.Once
- This also implements thread safety

## Problems

- This leads to a problem where we dont follow the `Dependency Inversion Principle`as directly depend on the data that is coming from the db.
- So, there would be no way to do unit tests here as it depends on the db being connected.
- But we can fix that by implementing an interface and interacting with that in any function.

```go
//  Directly depends on the Singleton DB
func GetTotalPopulation(cities []string) int {
 result := 0
 for _, city := range cities {
  result += GetSingletonDB().GetPopulation(city)
 }

 return result
}

//  Using the Datbase interface, we can add in a layer of abstraction
type Database interface {
 GetPopulation(name string) int
}
func GetTotalPopulationEx(db Database, cities []string) int {
 result := 0
 for _, city := range cities {
  result += db.GetPopulation(city)
 }

 return result
}

func UnitTest() {
 names := []string{"alpha", "gamma"}

 // dependency inversion
 tp := GetTotalPopulationEx(&DummyDatabase{}, names)
 fmt.Println(tp == 4)
}
```

- As you can see here, with the `Database` interface, we can solve the problem of hard dependencies which could ease us in writing unit tests.
