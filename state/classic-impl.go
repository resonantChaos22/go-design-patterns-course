package main

import "fmt"

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}
func (s *Switch) Off() {
	s.State.Off(s)
}

func NewSwitch() *Switch {
	return &Switch{
		State: &OffState{},
	}
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct{}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on!")
}
func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off!")
}

type OnState struct {
	BaseState
}

func (p *OnState) Off(sw *Switch) {
	fmt.Println("Turning the light off")
	sw.State = NewOffState()
}

func NewOnState() *OnState {
	fmt.Println("Light Turned On")
	return &OnState{
		BaseState: BaseState{},
	}
}

type OffState struct {
	BaseState
}

func (p *OffState) On(sw *Switch) {
	fmt.Println("Turning the light on")
	sw.State = NewOnState()
}

func NewOffState() *OffState {
	fmt.Println("Light Turned Off")
	return &OffState{
		BaseState: BaseState{},
	}
}

func TestClassicImplementation() {
	s := NewSwitch()
	s.On()
	s.Off()
	s.Off()
}
