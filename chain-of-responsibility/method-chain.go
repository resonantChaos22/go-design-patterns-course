package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{
		Name:    name,
		Attack:  attack,
		Defense: defense,
	}
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

// This struct is mainly used as a root modifier and for other modifiers to inherit Add() function
type CreatureModifier struct {
	creature *Creature
	next     Modifier //	it's a modifier not creature modifier
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

func NewCreatureModifier(c *Creature) *CreatureModifier {
	return &CreatureModifier{
		creature: c,
	}
}

// double attack modifier definition
type DoubleAttackModifier struct {
	CreatureModifier
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{
			creature: c,
		},
	}
}

// triple defense modifier definition
type TripleDefenseModifier struct {
	CreatureModifier
}

func (t *TripleDefenseModifier) Handle() {
	fmt.Println("Tripling", t.creature.Name, "\b's defense")
	t.creature.Defense *= 3
	t.CreatureModifier.Handle()
}

func NewTripleDefenseModifer(c *Creature) *TripleDefenseModifier {
	return &TripleDefenseModifier{
		CreatureModifier: CreatureModifier{
			creature: c,
		},
	}
}

// no bonuses modifier
type NoBonusModifier struct {
	CreatureModifier
}

func (n *NoBonusModifier) Handle() {
	//	empty
}

func NewNoBonusModifier(c *Creature) *NoBonusModifier {
	return &NoBonusModifier{
		CreatureModifier: CreatureModifier{
			creature: c,
		},
	}
}

func TestMethodChain() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewTripleDefenseModifer(goblin))
	root.Add(NewDoubleAttackModifier(goblin))

	root.Add(NewNoBonusModifier(goblin))
	root.Add(NewTripleDefenseModifer(goblin)) //	wont get applied, also no other modifiers can be added at all
	root.Handle()
	fmt.Println(goblin.String())
}

//	Note that there's a bug here.
//	Let's say we define a new creature "elf" and if we pass it as `root.Add(NewDoubleAttackModifer(elf))`,
//	there will be no errors thrown and there would be no changes in goblin
//	I think the modifier is defined on a global level and not on creature level, so at any point of time,
//	there would be tons of modifier in a singly linked list which will be specifying what bonuses which
//	all characters have
