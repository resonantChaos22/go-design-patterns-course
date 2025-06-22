package broker

import (
	"fmt"
	"sync"
)

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

// attack and defense are lowercase, to make it private, so that this does not change
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

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{
		game:    game,
		Name:    name,
		attack:  attack,
		defense: defense,
	}
}

func NewCreatureFluent() *Creature {
	return &Creature{}
}
func (c *Creature) WithGame(game *Game) *Creature {
	c.game = game
	return c
}
func (c *Creature) WithName(name string) *Creature {
	c.Name = name
	return c
}
func (c *Creature) WithAttack(attack int) *Creature {
	c.attack = attack
	return c
}
func (c *Creature) WithDefense(defense int) *Creature {
	c.defense = defense
	return c
}

func TestBrokerChain() {
	game := &Game{
		observers: sync.Map{},
	}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	elf := NewCreatureFluent().WithGame(game).WithName("Elf").WithAttack(4).WithDefense(6)
	fmt.Println(goblin.String())
	fmt.Println(elf.String())

	{
		m := NewDoubleAttackModifer(game, goblin)
		n := NewDoubleAttackModifer(game, goblin)
		o := NewDoubleAttackModifer(game, elf)
		fmt.Println(goblin.String())
		fmt.Println(elf.String())
		m.Close()
		n.Close()
		o.Close()
	}

	fmt.Println(goblin.String())
	fmt.Println(elf.String())

}
