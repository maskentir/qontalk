package fsm_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maskentir/qontalk/fsm"
)

func TestFSM_AddTransition(t *testing.T) {
	initialState := fsm.State("start")
	transitions := []fsm.Transition{
		{
			From:    fsm.State("start"),
			Event:   fsm.Event("event1"),
			To:      fsm.State("state1"),
			Action:  nil,
			OnError: nil,
			Timeout: 0,
		},
		{
			From:    fsm.State("state1"),
			Event:   fsm.Event("event2"),
			To:      fsm.State("state2"),
			Action:  nil,
			OnError: nil,
			Timeout: 0,
		},
	}

	fsmInstance, err := fsm.NewFSM(initialState, transitions, nil)
	if err != nil {
		t.Errorf("Failed to create FSM: %v", err)
	}

	testCases := []struct {
		from   fsm.State
		event  fsm.Event
		exists bool
	}{
		{fsm.State("start"), fsm.Event("event1"), true},
		{fsm.State("state1"), fsm.Event("event2"), true},
		{fsm.State("start"), fsm.Event("event2"), false},
		{fsm.State("state1"), fsm.Event("event1"), false},
		{fsm.State("state2"), fsm.Event("event3"), false},
	}

	for _, tc := range testCases {
		// Check if the transitions exist
		exists, _ := fsmInstance.TransitionExists(tc.from, tc.event)
		if tc.exists && !exists {
			t.Errorf("Expected transition to exist, but it does not exist: %v -> %v", tc.from, tc.event)
		} else if !tc.exists && exists {
			t.Errorf("Expected transition not to exist, but it exists: %v -> %v", tc.from, tc.event)
		}
	}
}

func TestFSM_RemoveTransition(t *testing.T) {
	initialState := fsm.State("start")
	transitions := []fsm.Transition{
		{
			From:    fsm.State("start"),
			Event:   fsm.Event("event1"),
			To:      fsm.State("state1"),
			Action:  nil,
			OnError: nil,
			Timeout: 0,
		},
		{
			From:    fsm.State("state1"),
			Event:   fsm.Event("event2"),
			To:      fsm.State("state2"),
			Action:  nil,
			OnError: nil,
			Timeout: 0,
		},
	}

	fsmInstance, err := fsm.NewFSM(initialState, transitions, nil)
	if err != nil {
		t.Errorf("Failed to create FSM: %v", err)
	}

	testCases := []struct {
		from   fsm.State
		event  fsm.Event
		exists bool
	}{
		{fsm.State("start"), fsm.Event("event1"), true},
		{fsm.State("state1"), fsm.Event("event2"), true},
		{fsm.State("start"), fsm.Event("event2"), false},
		{fsm.State("state1"), fsm.Event("event1"), false},
		{fsm.State("state2"), fsm.Event("event3"), false},
	}

	for _, tc := range testCases {
		// Remove the transition
		err := fsmInstance.RemoveTransition(tc.from, tc.event)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		// Check if the transition exists
		exists, _ := fsmInstance.TransitionExists(tc.from, tc.event)
		if tc.exists && exists {
			t.Errorf("Expected transition from %v to %v via %v to be removed, but it still exists", initialState, tc.from, tc.event)
		} else if !tc.exists && exists {
			t.Errorf("Expected transition from %v to %v via %v not to exist, but it still exists", initialState, tc.from, tc.event)
		}
	}
}

func TestFSM_RunWithTimeout(t *testing.T) {
	initialState := fsm.State("start")
	transitions := []fsm.Transition{
		{
			From:    fsm.State("start"),
			Event:   fsm.Event("event1"),
			To:      fsm.State("state1"),
			Action:  nil,
			OnError: nil,
			Timeout: 1, // Reduced timeout to 1 second
		},
	}

	fsmInstance, err := fsm.NewFSM(initialState, transitions, nil)
	if err != nil {
		t.Errorf("Failed to create FSM: %v", err)
	}

	// Test the transition with a timeout
	startTime := time.Now()
	err = fsmInstance.SendEvent(fsm.Event("event1"), nil)
	elapsedTime := time.Since(startTime)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if currentState := fsmInstance.GetCurrentState(); currentState != fsm.State("state1") {
		t.Errorf("Expected current state to be 'state1', but got: %v", currentState)
	}
	if elapsedTime.Seconds() > 1 {
		t.Errorf("Expected elapsed time to be less than 1 second, but got: %v seconds", elapsedTime.Seconds())
	}
}

func TestFSM_GlobalCallback(t *testing.T) {
	callbackCalled := 0
	callback := fsm.Callback(func(from fsm.State, event fsm.Event, to fsm.State, params map[string]interface{}) {
		fmt.Println("Callback called")
		callbackCalled++
	})

	fsmInstance, err := fsm.NewFSM(fsm.State("state1"), []fsm.Transition{
		{
			From:    fsm.State("state1"),
			Event:   fsm.Event("event1"),
			To:      fsm.State("state2"),
			Action:  func() error { return nil },
			OnError: fsm.State("errorState"),
		},
		// Transisi untuk kembali ke state1 dari state2
		{
			From:    fsm.State("state2"),
			Event:   fsm.Event("event2"),
			To:      fsm.State("state1"),
			Action:  func() error { return nil },
			OnError: fsm.State("errorState"),
		},
	}, callback)

	if err != nil {
		t.Fatal(err)
	}

	if err := fsmInstance.SendEvent(fsm.Event("event1"), nil); err != nil {
		t.Fatal(err)
	}

	// Sleep for a short time to allow transitions to complete.
	time.Sleep(100 * time.Millisecond)

	currentState := fsmInstance.GetCurrentState()
	if currentState != fsm.State("state2") {
		t.Errorf("Expected current state to be 'state2', but got '%v'", currentState)
	}

	if err := fsmInstance.SendEvent(fsm.Event("event2"), nil); err != nil {
		t.Fatal(err)
	}

	// Sleep for a short time to allow transitions to complete.
	time.Sleep(100 * time.Millisecond)

	currentState = fsmInstance.GetCurrentState()
	if currentState != fsm.State("state1") {
		t.Errorf("Expected current state to be 'state1', but got '%v'", currentState)
	}

	if callbackCalled != 2 {
		t.Log(currentState)
		t.Errorf("Expected global callback to be called 2 times, but got %d", callbackCalled)
	}
}

func TestFSM_Stop(t *testing.T) {
	initialState := fsm.State("start")
	transitions := []fsm.Transition{
		{
			From:    fsm.State("start"),
			Event:   fsm.Event("event1"),
			To:      fsm.State("state1"),
			Action:  nil,
			OnError: nil,
			Timeout: 0,
		},
		{
			From:    fsm.State("state1"),
			Event:   fsm.Event("event2"),
			To:      fsm.State("state2"),
			Action:  nil,
			OnError: nil,
			Timeout: 0,
		},
	}

	fsmInstance, err := fsm.NewFSM(initialState, transitions, nil)
	if err != nil {
		t.Errorf("Failed to create FSM: %v", err)
	}

	err = fsmInstance.SendEvent(fsm.Event("event1"), nil)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	err = fsmInstance.SendEvent(fsm.Event("event2"), nil)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	go func() {
		time.Sleep(time.Second)
		fsmInstance.Stop()
	}()

	startTime := time.Now()
	err = fsmInstance.SendEvent(fsm.Event("event2"), nil)
	elapsedTime := time.Since(startTime)

	if err == nil {
		t.Errorf("Expected an error after stopping FSM, but got none")
	}
	if currentState := fsmInstance.GetCurrentState(); currentState != fsm.State("state2") {
		t.Errorf("Expected current state to remain 'state2' after stopping FSM, but got: %v", currentState)
	}
	if elapsedTime.Seconds() > 1 {
		t.Errorf("Expected elapsed time to be less than 1 second after stopping FSM, but got: %v seconds", elapsedTime.Seconds())
	}
}
