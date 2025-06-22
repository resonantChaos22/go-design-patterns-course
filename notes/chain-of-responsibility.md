- The idea is a chain of components who all get a chance to process a command or query, optionally having default processing implementation and an ability to terminate the processing chain.
- Can be implemented as a linked list o pointers or a centralised construct (like sync.Map in the second example)
- Enlist objects in chain, maybe control their order
- Control Object removal from chain.

## Command Query Separation

- `Command` - do modification to the data
- `Query` - get the data
- CQS (Command Query Separation) - having separate means of sending commands and queries.

## Creature Game

- We have a `Creature` struct to which we can add modifiers to update its stats like Defense and Attack

```go
type Modifier interface {
 Add(m Modifier)
 Handle()
}

// This struct is mainly used as a root modifier and for other modifiers to inherit Add() function
type CreatureModifier struct {
 creature *Creature
 next     Modifier // it's a modifier not creature modifier
}

// it's a modifier not creature modifier
func (c *CreatureModifier) Add(m Modifier) {
 if c.next != nil {
  c.next.Add(m)
 } else {
  c.next = m
 }
}
func (c *CreatureModifier) Handle() {
 if c.next != nil {
  c.next.Handle()
 }
}

type DoubleAttackModifier struct {
 CreatureModifier
}

func (d *DoubleAttackModifier) Handle() {
 fmt.Println("Doubling", d.creature.Name, "\b's attack")
 d.creature.Attack *= 2
 d.CreatureModifier.Handle()
}

root := NewCreatureModifier(goblin)
root.Add(NewDoubleAttackModifier(goblin))
root.Handle()
```

- Here, `Modifier` is an interface about how to define any Modifier struct, this is especially useful in understanding how `Add` and `CreatureModifier` is defined.
- `CreatureModifier` is a base struct which gives the base functionality to all Modifiers
- Both its `Add` and `Handle` functions work upon a Singly-Linked List of structs of `Modifier` interface.
- `Add` basically creates a chain of Modifiers
- `Handle` goes through the chain of modifiers and executes their `Handle()` function to update the `Creature` object based on different implementations.
- As you can see in the implementation, `root` is of type `CreatureModifier` and serves as the entry point for all the modifiers.
- This is a `Chain Of Responsibility` pattern as on each step of the linked list of modifiers, we can choose or not choose to do any changes.
- Essentially, we are setting up a pipeline of "bonuses" or "modifiers" for a `Creature`, and the `CreatureModifier` and its concrete implementations dictate how that request for modification flows through the pipeline.

## Broker Chain

- Another way of implementing the `Creature` using Observers and Command Query Separation
- The basic idea is that the attack and defense of the Creature dont change but the new attack and defense are calculated based on queries.
- First, let's talk about `Observer` and `Game`

```go
type Argument int

const (
 Attack Argument = iota
 Defense
)

// the `Value` can be initial value but it can also be modified value after going through modifiers
type Query struct {
 CreatureName string
 WhatToQuery  Argument
 Value        int
}

// basically all the modifiers
type Observer interface {
 Handle(q *Query)
}

// to handle the queries and do all the changes based on the modifiers
type Observable interface {
 Subscribe(o Observer)
 Unsubscribe(o Observer)
 Fire(q *Query)
}

type Game struct {
 observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
 g.observers.Store(o, struct{}{})
}
func (g *Game) Unsubscribe(o Observer) {
 g.observers.Delete(o)
}
func (g *Game) Fire(q *Query) {
 g.observers.Range(func(key, value any) bool {
  if key == nil {
   return false
  }
  key.(Observer).Handle(q)
  return true
 })
}
```

- So, let's understand with a dry run. When `Fire()` method is called with a `Query` object, it basically calls `Handle()` function of each observer that is subscribed to it.
- For our understanding, we can say that here `Observer` is nothing but `Modifier`, so on calling `Fire()`, the initial values of attack and defense go through each `Observer` (Modifier)'s `Handle()` function and are updated accordingly. This follows the `Chain Of Responsibily` pattern.
- The `Fire()` function follows the `Observer` pattern as it relays the information to all of the `Observer`s subscribed to it.
- Also, in this case, there are no Commands, only Queries as the internal value of attack and defense never change.

```go
type Creature struct {
 game            *Game
 Name            string
 attack, defense int
}

func (c *Creature) Attack() int {
 q := Query{
  CreatureName: c.Name,
  WhatToQuery:  Attack,
  Value:        c.attack,
 }
 c.game.Fire(&q)
 return q.Value
}
func (c *Creature) Defense() int {
 q := Query{
  CreatureName: c.Name,
  WhatToQuery:  Defense,
  Value:        c.defense,
 }
 c.game.Fire(&q)
 return q.Value
}
func (c *Creature) String() string {
 return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}
```

- The `Fire()` method is used to get the attack and defense. We send the initial attack as the value and the `Value` is updated over time as it passes through all the `Handle()` functions of the observers (modifiers)
- And here's how we implement the Modifiers which follow the `Observer` interface

```go
type CreatureModifier struct {
 game     *Game
 creature *Creature
}

func (c *CreatureModifier) Handle(q *Query) {

}

type DoubleAttackModifier struct {
 CreatureModifier
}

func (d *DoubleAttackModifier) Handle(q *Query) {
 if q.CreatureName == d.creature.Name && q.WhatToQuery == Attack {
  q.Value *= 2
 }
}
func (d *DoubleAttackModifier) Close() error {
 d.game.Unsubscribe(d)
 return nil
}

// here if we pass creature Modifier directly, we can make sure the changes are only being done for one creature
func NewDoubleAttackModifer(g *Game, c *Creature) *DoubleAttackModifier {
 d := &DoubleAttackModifier{
  CreatureModifier: CreatureModifier{
   game:     g,
   creature: c,
  },
 }
 g.Subscribe(d)
 return d
}
```

- As we can see here, the `Handle()` method of this identifier doubles the query's value if the query is about `Attack`
- In the factory function to initialize the Modifier, we `Subscribe` to the game, and hence add to the `sync.Map` of Observers that process the query.
- We have a `Close()` function if we want to remove any modifier which just unsubscribes.
