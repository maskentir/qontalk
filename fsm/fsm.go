// Package fsm provides a Finite State Machine (FSM) implementation in Go.
//
// # Overview
//
// The fsm package allows you to create and manage a Finite State Machine (FSM).
// It defines types and methods to represent states, transitions, rules, and actions
// within the FSM. You can create an FSM, define states, transitions between states,
// and rules to handle user input. The FSM processes user messages and provides
// responses based on the current state, transitions, and defined rules.
//
// # Bot
//
// The Bot struct represents the FSM-based chatbot. It allows you to create and manage
// a chatbot instance with multiple states, rules, and actions.
//
// # FsmState
//
// The FsmState struct represents a state within the FSM. It defines the state's name,
// entry message, transitions to other states, rules to handle messages, and a custom error rule.
//
// # Transition
//
// The Transition struct defines a state transition triggered by a specific event. It specifies
// the event name and the target state after the transition.
//
// # Rule
//
// The Rule struct represents a rule for handling user messages within a state. It defines
// a regular expression pattern to match user input, a response message template, and actions
// to perform when the rule is triggered.
//
// # Action
//
// The Action struct represents an action to be performed when a rule is triggered. Currently,
// the only action type supported is SetVariableAction.
//
// # SetVariableAction
//
// The SetVariableAction struct represents an action that sets a variable's value in the user's session.
// It allows you to store and manipulate data during the conversation.
//
// # UserSession
//
// The UserSession struct represents a user's session with the chatbot. It stores session variables
// and the current session state.
//
// # Getting Started
//
// To create and use the chatbot FSM:
// 1. Create a new bot instance with NewBot.
// 2. Add states using AddState, specifying their name, entry message, transitions, rules, and custom error rules.
// 3. Add rules to states using AddRuleToState, defining regular expressions and responses.
// 4. Process user messages with ProcessMessage, which handles state transitions and rule execution.
//
// # Example
//
// Here's an example of how to use the fsm package to create and use a chatbot FSM:
//
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/yourusername/fsm"
//	    "regexp"
//	    "time"
//	)
//
//	func main() {
//	    // Create a new chatbot instance
//	    bot := fsm.NewBot("MyChatbot")
//
//	    // Define states and transitions
//	    transitions := []fsm.Transition{
//	        {Event: "start", Target: "initial"},
//	        {Event: "continue", Target: "ongoing"},
//	    }
//
//	    bot.AddState("initial", "Welcome to the chatbot!", transitions)
//
//	    // Define rules and actions
//	    rulePattern := "hello"
//	    regexPattern := fmt.Sprintf("(?i)%s", regexp.QuoteMeta(rulePattern))
//	    rule := fsm.Rule{
//	        Name:    "HelloRule",
//	        Pattern: regexp.MustCompile(regexPattern),
//	        Respond: "Hello! How can I assist you?",
//	    }
//
//	    bot.AddRuleToState("initial", rule.Name, regexPattern, rule.Respond, nil, nil)
//
//	    // Process user messages
//	    response, err := bot.ProcessMessage("user123", "hello")
//	    if err != nil {
//	        fmt.Println("Error:", err)
//	        return
//	    }
//
//	    fmt.Println("Bot Response:", response)
//	}
//
// For more information on how to use this FSM-based chatbot, please refer to the package documentation.
package fsm

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Bot represents the FSM-based chatbot.
type Bot struct {
	Name             string
	CurrentState     string
	UserSessions     map[string]*UserSession
	UserMutex        sync.RWMutex
	FsmStates        map[string]*FsmState
	GlobalVars       map[string]string
	StateListeners   map[string]ListenerFunc
	RuleListeners    map[string]ListenerFunc
	SessionTimeout   time.Duration
	SessionCleanup   time.Duration
	ConcurrentAccess bool
	ErrorLogger      func(error)
	stopCleanup      chan struct{}
}

// FsmState represents a state within the FSM.
type FsmState struct {
	Name         string
	EntryMessage string
	Transitions  []Transition
	Rules        []Rule
}

// Transition defines a state transition in the FSM.
type Transition struct {
	Event  string
	Target string
}

// CustomError represents a custom error rule for handling specific errors.
type CustomError struct {
	Error   error
	Respond string
}

// Rule represents a rule for handling user messages within a state.
type Rule struct {
	Name       string
	Pattern    *regexp.Regexp
	Respond    string
	Actions    []Action
	ErrorRules []CustomError
}

// Action represents an action to be performed when a rule is triggered.
type Action struct {
	SetVariable *SetVariableAction
}

// SetVariableAction represents an action that sets a variable's value in the user's session.
type SetVariableAction struct {
	Name  string
	Value string
}

// VariableMap is a type alias for a map of string variables.
type VariableMap map[string]string

// ListenerFunc represents a listener function.
type ListenerFunc func(userID string, message string, session *UserSession, bot *Bot)

// UserSession represents a user's session with the chatbot.
type UserSession struct {
	// SessionVars is a map of session variables.
	SessionVars VariableMap

	// SessionState is the current state of the user's session.
	SessionState string

	// LastActive is the timestamp when the user was last active.
	LastActive time.Time

	// ErrorRulesState is a map of error rules associated with each state.
	ErrorRulesState map[string]map[string]bool

	// ErrorRulesChan is a channel for updating error rules state.
	ErrorRulesChan chan map[string]map[string]bool
}

// cleanupSessions periodically cleans up inactive user sessions.
func (b *Bot) cleanupSessions() {
	for {
		select {
		case <-time.After(b.SessionCleanup):
			b.UserMutex.Lock()
			for userID, session := range b.UserSessions {
				if time.Since(session.LastActive) > b.SessionTimeout {
					delete(b.UserSessions, userID)
				}
			}
			b.UserMutex.Unlock()
		case <-b.stopCleanup:
			return
		}
	}
}

// NewBot creates a new chatbot instance with the specified name and options.
func NewBot(name string, options ...Option) *Bot {
	bot := &Bot{
		Name:             name,
		CurrentState:     "start",
		UserSessions:     make(map[string]*UserSession),
		FsmStates:        make(map[string]*FsmState),
		GlobalVars:       make(map[string]string),
		StateListeners:   make(map[string]ListenerFunc),
		RuleListeners:    make(map[string]ListenerFunc),
		SessionTimeout:   30 * time.Minute,
		SessionCleanup:   1 * time.Hour,
		ConcurrentAccess: false,
		ErrorLogger:      nil,
		stopCleanup:      make(chan struct{}),
	}

	for _, option := range options {
		option(bot)
	}

	if bot.SessionCleanup > 0 {
		go bot.cleanupSessions()
	}

	return bot
}

// Option represents an option to configure the chatbot.
type Option func(*Bot)

// WithSessionCleanup sets the session cleanup interval for removing inactive sessions.
func WithSessionCleanup(interval time.Duration) Option {
	return func(b *Bot) {
		b.SessionCleanup = interval
	}
}

// WithSessionTimeout sets the session timeout interval for removing inactive sessions.
func WithSessionTimeout(interval time.Duration) Option {
	return func(b *Bot) {
		b.SessionTimeout = interval
	}
}

// WithConcurrentAccess enables or disables concurrent access handling.
func WithConcurrentAccess(enable bool) Option {
	return func(b *Bot) {
		b.ConcurrentAccess = enable
	}
}

// WithErrorLogger sets the error logger function for handling errors.
func WithErrorLogger(logger func(error)) Option {
	return func(b *Bot) {
		b.ErrorLogger = logger
	}
}

// AddState adds a state to the chatbot's FSM.
func (b *Bot) AddState(name, entryMessage string, transitions []Transition) {
	state := &FsmState{
		Name:         name,
		EntryMessage: entryMessage,
		Transitions:  transitions,
	}
	b.FsmStates[name] = state
}

// AddRuleToState adds a rule to a specific state.
func (b *Bot) AddRuleToState(stateName, name, pattern, respond string, actions []Action, errorRules []CustomError) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	rule := Rule{
		Name:    name,
		Pattern: re,
		Respond: respond,
		Actions: actions,
	}

	if errorRules != nil {
		rule.ErrorRules = errorRules
	}

	state, ok := b.FsmStates[stateName]
	if !ok {
		return fmt.Errorf("state %s not found", stateName)
	}

	state.Rules = append(state.Rules, rule)
	b.FsmStates[stateName] = state
	return nil
}

// AddListenerToState adds a listener function to a specific state.
func (b *Bot) AddListenerToState(stateName string, listener ListenerFunc) {
	b.StateListeners[stateName] = listener
}

// AddListenerToRule adds a listener function to a specific rule.
func (b *Bot) AddListenerToRule(ruleName string, listener ListenerFunc) {
	b.RuleListeners[ruleName] = listener
}

// ProcessMessage processes a user's message and returns a response based on the chatbot's current state.
func (b *Bot) ProcessMessage(userID, message string) (string, error) {
	b.UserMutex.Lock()
	defer b.UserMutex.Unlock()

	session, ok := b.UserSessions[userID]
	if !ok {
		session = &UserSession{
			SessionVars:  make(VariableMap),
			SessionState: b.CurrentState,
		}
		b.UserSessions[userID] = session
	}

	session.LastActive = time.Now()
	state, ok := b.FsmStates[session.SessionState]
	if !ok {
		b.handleError("State not found", userID, session)
		return "State not found", nil
	}

	if session.ErrorRulesChan == nil {
		session.ErrorRulesChan = make(chan map[string]map[string]bool)
	}

	stopErrorRules := make(chan struct{})
	defer close(stopErrorRules)

	go func() {
		for {
			select {
			case updatedErrorRules := <-session.ErrorRulesChan:
				if updatedErrorRules != nil {
					session.ErrorRulesState = updatedErrorRules
				}
			case <-stopErrorRules:
				return
			}
		}
	}()

	for _, transition := range state.Transitions {
		if transition.Event == message {
			if transition.Target == "start" {
				session.SessionState = "start"
			} else {
				session.SessionState = transition.Target
			}
			b.CurrentState = session.SessionState
			state = b.FsmStates[b.CurrentState] // Update state to the new one
			entryMessage := b.replaceVariables(state.EntryMessage, session.SessionVars)
			b.handleStateListener(state.Name, userID, message, session)
			return entryMessage, nil
		}
	}

	var (
		wg        sync.WaitGroup
		respChan  = make(chan string, len(state.Rules))
		errorChan = make(chan error, len(state.Rules))
	)

	foundValidRule := false

	for _, rule := range state.Rules {
		wg.Add(1)

		go func(rule Rule) {
			defer wg.Done()

			match := rule.Pattern.FindStringSubmatch(message)
			if match != nil {
				foundValidRule = true

				for i, name := range rule.Pattern.SubexpNames() {
					if i > 0 && name != "" {
						session.SessionVars[name] = match[i]
					}
				}

				for _, action := range rule.Actions {
					if action.SetVariable != nil {
						if value, ok := session.SessionVars[action.SetVariable.Value]; ok {
							session.SessionVars[action.SetVariable.Name] = value
						}
					}
				}

				respond := rule.Respond
				respond = b.replaceVariables(respond, session.SessionVars)

				b.handleStateListener(state.Name, userID, message, session)
				b.handleRuleListener(rule.Name, userID, message, session)

				for _, errorRule := range rule.ErrorRules {
					if session.ErrorRulesState != nil && session.ErrorRulesState[state.Name][errorRule.Error.Error()] {
						b.handleError(errorRule.Respond, userID, session)
						respChan <- errorRule.Respond

						delete(session.ErrorRulesState, state.Name)
						return
					}

				}

				respChan <- respond
			}
		}(rule)
	}

	go func() {
		wg.Wait()
		close(respChan)
		close(errorChan)
	}()

	var responses []string
	for response := range respChan {
		responses = append(responses, response)
	}

	if len(responses) > 0 {
		return responses[len(responses)-1], nil
	}

	if !foundValidRule {
		b.handleError("No valid rule found", userID, session)
	}

	entryMessage := b.replaceVariables(state.EntryMessage, session.SessionVars)
	b.handleStateListener(state.Name, userID, message, session)
	return entryMessage, nil
}

// ProcessError processes an error associated with a specific rule in a state.
func (b *Bot) ProcessError(userID, stateName, ruleName string, err error) {
	currentState, ok := b.FsmStates[stateName]
	if !ok {
		return
	}

	for _, currentRule := range currentState.Rules {
		if currentRule.Name == ruleName {
			session, ok := b.UserSessions[userID]
			if !ok {
				return
			}

			if session.ErrorRulesState == nil {
				session.ErrorRulesState = make(map[string]map[string]bool)
			}

			if _, ok := session.ErrorRulesState[stateName]; !ok {
				session.ErrorRulesState[stateName] = make(map[string]bool)
			}

			session.ErrorRulesState[stateName][err.Error()] = true

			if session.ErrorRulesChan != nil {
				session.ErrorRulesChan <- session.ErrorRulesState
			}
		}
	}
}

// replaceVariables replaces variables in the text with their session values and global variables.
func (b *Bot) replaceVariables(text string, vars VariableMap) string {
	for name, value := range vars {
		placeholder := fmt.Sprintf("{{%s}}", name)
		text = strings.ReplaceAll(text, placeholder, value)
	}

	for name, value := range b.GlobalVars {
		placeholder := fmt.Sprintf("{{bot.%s}}", name)
		text = strings.ReplaceAll(text, placeholder, value)
	}

	return text
}

// handleStateListener calls the state listener function if available.
func (b *Bot) handleStateListener(stateName, userID, message string, session *UserSession) {
	if listener, ok := b.StateListeners[stateName]; ok {
		listener(userID, message, session, b)
	}
}

// handleRuleListener calls the rule listener function if available.
func (b *Bot) handleRuleListener(ruleName, userID, message string, session *UserSession) {
	if listener, ok := b.RuleListeners[ruleName]; ok {
		listener(userID, message, session, b)
	}
}

// handleError handles an error message by logging it and potentially notifying the user.
func (b *Bot) handleError(errorMessage, userID string, session *UserSession) {
	if b.ErrorLogger != nil {
		err := fmt.Errorf("error for user %s: %s", userID, errorMessage)
		b.ErrorLogger(err)
	}
}

// Stop stops the session cleanup goroutine.
func (b *Bot) Stop() {
	close(b.stopCleanup)
}
