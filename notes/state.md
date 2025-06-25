- A pattern in which the object's behaviour is determined by its state. An object transitions from one state to another because of a `trigger`
- A formalised construct which manages state and transitions is called a `state machine`
- We can define - 
	- State entry / exit behaviors
	- Action where a particular event causes a transition
	- Guard Conditions enabling / disabling a transition
	- Default actions when no transitions are found for an event.

## Classic Implementation

- In the classic implementation, states changes based on replacement and the method of the current state is responsible  for replacing the state.
- Let's take a `Switch` struct which has a `State`

```go
type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}
func (s *Switch) Off() {
	s.State.Off(s)
}

//  State and a Base State
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

//  On State implementation
type OnState struct {
	BaseState
}

func (p *OnState) Off(sw *Switch) {
	fmt.Println("Turning the light off")
	sw.State = NewOffState()
}

```

- `OffState` is symmetric to `OnState`, so basically the `Off()` method of `OnState` changes the state of the switch to be `OffState`
- The `On()` method is inherited from its parent.
- There is no `State Machine 
- There is no need to define state with expensive structs, we can do it with a simpler implementation.

## Handmade State Machine

- We work here with `State`, `Trigger` and `TriggerResult`
- We also work here with a map of `State` -> []`TriggerResult`
- The idea is `State` -> `Trigger` -> `Final State`
- The idea of `TriggerResult` is used weirdly in this example but we can utilise `Trigger` to switch from `State` to `FinalState`
```go
type TriggerResult struct {
	Trigger Trigger
	State   State
}

var rules = map[State][]TriggerResult{
	OffHook: {
		{Trigger: CallDialed, State: Connecting},
	},
	Connecting: {
		{Trigger: HungUp, State: OnHook},
		{Trigger: CallConnected, State: Connected},
	},
	Connected: {
		{Trigger: LeftMessage, State: OnHook},
		{Trigger: HungUp, State: OnHook},
		{Trigger: PlacedOnHold, State: OnHold},
	},
	OnHold: {
		{Trigger: TakenOffHold, State: Connected},
		{Trigger: HungUp, State: OnHook},
	},
}

func TestStateMachine() {
	state, exitState := OffHook, OnHook

	for ok := true; ok; ok = state != exitState {
		fmt.Println("The phone is currently", state)
		fmt.Println("Select a trigger: ")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}
	fmt.Println("Call Ended")
}

```
- As we can see here, it works by showing the output based on the selected trigger as at every state, there are some possibilities of the next state based on the available `triggers`

## Switch Based State Machine

- Switch based State machines use switch cases to switch the state.

```go
func TestSwitchBasedStateMachine() {
	code := "1234"
	state := Locked
	entry := new(strings.Builder)

	for {
		switch state {
		case Locked:
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			if entry.String() == code {
				state = Unlocked
				break
			}

			if strings.Index(code, entry.String()) != 0 {
				state = Failed
			}

		case Failed:
			fmt.Println("FAILED")
			entry.Reset()
			state = Locked

		case Unlocked:
			fmt.Println("UNLOCKED")
			return
		}
	}
}
```

- This was pretty fun