// Package fsm provides a finite state machine (FSM) implementation in Go.
//
// # Overview
//
// This package allows you to create and manage a finite state machine. It
// defines types and methods to represent states, events, and transitions
// within the FSM. You can create an FSM, define transitions between states
// triggered by events, and execute those transitions.
//
// # State
//
// The State type represents states within the FSM. States can be of any type
// as long as they are unique.
//
// # Event
//
// The Event type represents events that can trigger transitions in the FSM.
// Events can also be of any type, but they should be unique.
//
// # Callback
//
// Callback is a function type that can be defined to execute custom logic
// when transitioning from one state to another. It receives information about
// the previous state, the triggering event, the new state, and optional
// parameters.
//
// # Transition
//
// The Transition struct defines a state transition in the FSM. It specifies
// the initial state (From), the event that triggers the transition (Event),
// the target state after the transition (To), an optional action to execute
// during the transition (Action), an optional state to transition to in case
// of an error (OnError), and an optional timeout for the transition action.
//
// # FSM
//
// The FSM struct represents the finite state machine itself. You can create
// an FSM instance using NewFSM, define transitions with AddTransition and
// RemoveTransition, and send events to trigger transitions using SendEvent.
// You can also retrieve the current state with GetCurrentState and stop the
// FSM with Stop.
//
// # Example
//
// Below is an example of how to use the fsm package to create a simple FSM:
//
//	// Define custom state types.
//	type MyState int
//	const (
//	    StateA MyState = iota
//	    StateB
//	    StateC
//	)
//
//	// Define custom event types.
//	type MyEvent int
//	const (
//	    EventX MyEvent = iota
//	    EventY
//	)
//
//	// Define a callback function to execute when transitioning.
//	callback := func(from fsm.State, event fsm.Event, to fsm.State, params map[string]interface{}) {
//	    fmt.Printf("Transition from %v to %v due to event %v\n", from, to, event)
//	}
//
//	// Create an FSM instance with an initial state, transitions, and the callback.
//	transitions := []fsm.Transition{
//	    {From: StateA, Event: EventX, To: StateB},
//	    {From: StateB, Event: EventY, To: StateC},
//	}
//
//	fsmInstance, err := fsm.NewFSM(StateA, transitions, callback)
//	if err != nil {
//	    fmt.Println("Error creating FSM:", err)
//	    return
//	}
//
//	// Send events to trigger transitions.
//	fsmInstance.SendEvent(EventX, nil)
//	fsmInstance.SendEvent(EventY, nil)
//
//	// Get the current state.
//	currentState := fsmInstance.GetCurrentState()
//
//	// Stop the FSM.
//	fsmInstance.Stop()
package fsm

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// State represents a state in the finite state machine.
type State interface{}

// Event represents an event that can trigger a state transition.
type Event interface{}

// Callback is a function type that is called on state transitions. It receives
// the previous state, the triggering event, the new state, and optional parameters.
type Callback func(from State, event Event, to State, params map[string]interface{})

// Transition represents a state transition in the finite state machine.
type Transition struct {
	From    State         // The initial state of the transition.
	Event   Event         // The event that triggers the transition.
	To      State         // The target state after the transition.
	Action  func() error  // Optional action to execute during the transition.
	OnError State         // Optional state to transition to in case of an error.
	Timeout time.Duration // Optional timeout for the transition action.
}

// FSM represents a finite state machine.
type FSM struct {
	currentState   State                          // The current state of the FSM.
	transitions    map[State]map[Event]Transition // Mapping of states and events to transitions.
	globalCallback Callback                       // Global callback for all state transitions.
	mutex          sync.Mutex                     // Mutex for thread safety.
	stopCh         chan struct{}                  // Channel to stop the FSM.
	wg             sync.WaitGroup                 // WaitGroup for tracking goroutines.
}

// NewFSM creates a new FSM with the given initial state, transitions, and global callback.
func NewFSM(initialState State, transitions []Transition, globalCallback Callback) (*FSM, error) {
	// Validate the globalCallback parameter.
	if globalCallback == nil {
		return nil, errors.New("globalCallback cannot be nil")
	}

	// Create a new FSM instance.
	fsm := &FSM{
		currentState:   initialState,
		transitions:    make(map[State]map[Event]Transition),
		globalCallback: globalCallback,
		stopCh:         make(chan struct{}),
	}

	// Initialize the FSM with the provided transitions.
	for _, t := range transitions {
		if _, exists := fsm.transitions[t.From]; !exists {
			fsm.transitions[t.From] = make(map[Event]Transition)
		}

		if _, exists := fsm.transitions[t.From][t.Event]; exists {
			return nil, fmt.Errorf("duplicate transition: state %v already has a transition for event %v", t.From, t.Event)
		}

		if _, exists := fsm.transitions[t.To]; !exists {
			fsm.transitions[t.To] = make(map[Event]Transition)
		}

		fsm.transitions[t.From][t.Event] = t
	}

	return fsm, nil
}

// TransitionExists checks if a transition from a specified state to a specified event exists.
func (f *FSM) TransitionExists(from State, event Event) (bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	// Check if the from state exists in transitions.
	fromTransitions, ok := f.transitions[from]
	if !ok {
		return false, fmt.Errorf("transition from state '%s' does not exist", from)
	}

	// Check if the event exists in fromTransitions.
	_, exists := fromTransitions[event]

	return exists, nil
}

// SendEvent sends an event to the FSM, triggering a state transition if a valid transition exists.
func (f *FSM) SendEvent(event Event, params map[string]interface{}) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	stateMap, exists := f.transitions[f.currentState]
	if !exists {
		return errors.New("invalid state")
	}

	transition, valid := stateMap[event]
	if !valid {
		return errors.New("invalid event for the current state")
	}

	if err := f.validateTransition(transition); err != nil {
		return err
	}

	if transition.Timeout > 0 {
		f.runWithTimeout(transition, params)
	} else if transition.Action != nil {
		// Run the transition action in a goroutine.
		go f.runWithAction(transition, params)
	}

	if f.currentState == transition.From {
		// Transition to the new state.
		f.currentState = transition.To

		// Log the current state after the transition.
		fmt.Printf("Current state after transition: %v\n", f.currentState)

		return nil
	}

	return errors.New("invalid transition")
}

// GetCurrentState returns the current state of the FSM.
func (f *FSM) GetCurrentState() State {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.currentState
}

// Stop stops the FSM and waits for all goroutines to complete.
func (f *FSM) Stop() {
	close(f.stopCh)
	f.wg.Wait()
}

// AddTransition adds a new transition to the FSM.
func (f *FSM) AddTransition(t Transition) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if _, exists := f.transitions[t.From]; !exists {
		f.transitions[t.From] = make(map[Event]Transition)
	}
	if _, exists := f.transitions[t.To]; !exists || t.From == t.To {
		return nil
	}

	f.transitions[t.From][t.Event] = t
	return nil
}

// RemoveTransition removes a transition from the FSM.
func (f *FSM) RemoveTransition(from State, event Event) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	stateMap, exists := f.transitions[from]
	if exists {
		delete(stateMap, event)
		return nil
	}
	return errors.New("invalid state")
}

// runWithTimeout runs a transition action with a specified timeout.
func (f *FSM) runWithTimeout(transition Transition, params map[string]interface{}) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()

		select {
		case <-time.After(transition.Timeout):
			f.mutex.Lock()
			currentState := f.currentState
			defer f.mutex.Unlock()

			if currentState == transition.From {
				err := transition.Action()
				if err != nil {
					f.handleTransitionError(transition, err, params)
				} else {
					// Update currentState only if it is still the same as transition.From
					f.mutex.Lock()
					if f.currentState == transition.From {
						f.currentState = transition.To
					}
					defer f.mutex.Unlock()
					f.globalCallback(currentState, transition.Event, transition.To, params)
				}
			}
		case <-f.stopCh:
			return
		}
	}()
}

// runWithAction runs a transition action without a timeout.
func (f *FSM) runWithAction(transition Transition, params map[string]interface{}) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()

		err := transition.Action()
		if err != nil {
			f.mutex.Lock()
			defer f.mutex.Unlock()

			if f.currentState == transition.From {
				f.handleTransitionError(transition, err, params)
			}
		} else {
			f.mutex.Lock()
			if f.currentState == transition.From {
				f.currentState = transition.To
			}
			defer f.mutex.Unlock()

			f.globalCallback(transition.From, transition.Event, transition.To, params)
		}
	}()
}

// handleTransitionError handles a transition error, including transitioning to an error state and invoking the global callback.
func (f *FSM) handleTransitionError(transition Transition, err error, params map[string]interface{}) {
	if transition.OnError != nil {
		f.currentState = transition.OnError
	}
	f.globalCallback(transition.From, transition.Event, transition.OnError, params)
}

// validateTransition validates whether a transition is valid for the current state and event.
func (f *FSM) validateTransition(transition Transition) error {
	stateMap, exists := f.transitions[f.currentState]
	if !exists {
		return errors.New("invalid state")
	}

	if _, valid := stateMap[transition.Event]; !valid {
		return errors.New("invalid event for the current state")
	}

	return nil
}
