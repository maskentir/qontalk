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
//		"fmt"
//		"fsm" // Replace with the actual package path.
//	)
//
//	func main() {
//		bot := fsm.NewBot("ChatBot")
//
//		bot.AddState("start", "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", []fsm.Transition{
//			{Event: "1", Target: "view_growth_history"},
//			{Event: "2", Target: "update_growth_data"},
//		}, []fsm.Rule{}, fsm.Rule{})
//
//		bot.AddState("view_growth_history", "Growth history of your child: Name: {{child_name}} Height: {{height}} Weight: {{weight}} Month: {{month}}", []fsm.Transition{
//			{Event: "exit", Target: "start"},
//		}, []fsm.Rule{}, fsm.Rule{
//			Name:    "custom_error",
//			Pattern: regexp.MustCompile("error"),
//			Respond: "Custom error message for view_growth_history state.",
//		})
//
//		bot.AddState("update_growth_data", "Please provide the growth information for your child. Use this template e.g., 'Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm'", []fsm.Transition{
//			{Event: "exit", Target: "start"},
//		}, []fsm.Rule{}, fsm.Rule{
//			Name:    "custom_error",
//			Pattern: regexp.MustCompile("error"),
//			Respond: "Custom error message for update_growth_data state.",
//		})
//
//		bot.AddRuleToState("update_growth_data", "rule_update_growth_data", `Month: (?P<month>.+) Child's name: (?P<child_name>.+) Weight: (?P<weight>.+) kg Height: (?P<height>.+) cm`, "Thank you for updating {{child_name}}'s growth in {{month}} with height {{height}} and weight {{weight}}", nil)
//
//		messages := []string{
//			"2",
//			"Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm",
//			"error",
//		}
//
//		for _, message := range messages {
//			response, err := bot.ProcessMessage("user1", message)
//			if err != nil {
//				fmt.Printf("Error processing message '%s': %v\n", message, err)
//			} else {
//				fmt.Printf("User1: %s\n", message)
//				fmt.Printf("Bot: %s\n", response)
//			}
//		}
//	}
//
// Bot represents a finite state machine (FSM) bot.
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
	UserMutex        sync.Mutex
	FsmStates        map[string]*FsmState
	GlobalVars       map[string]string
	StateListeners   map[string]ListenerFunc
	RuleListeners    map[string]ListenerFunc
	SessionTimeout   time.Duration
	SessionCleanup   time.Duration
	ConcurrentAccess bool
	ErrorLogger      func(error)
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

// Rule represents a rule for handling user messages within a state.
type Rule struct {
	Name      string
	Pattern   *regexp.Regexp
	Respond   string
	Actions   []Action
	ErrorRule *ErrorRule
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

// ErrorRule represents an error rule for handling specific errors.
type ErrorRule struct {
	Name    string
	Pattern *regexp.Regexp
	Respond string
	Actions []Action
}

// VariableMap is a type alias for a map of string variables.
type VariableMap map[string]string

// ListenerFunc represents a listener function.
type ListenerFunc func(userID string, message string, session *UserSession, bot *Bot)

// UserSession represents a user's session with the chatbot.
type UserSession struct {
	SessionVars  VariableMap
	SessionState string
	LastActive   time.Time
}

// cleanupSessions periodically cleans up inactive user sessions.
func (b *Bot) cleanupSessions() {
	for {
		time.Sleep(b.SessionCleanup)

		b.UserMutex.Lock()
		for userID, session := range b.UserSessions {
			fmt.Println("since", session.LastActive, b.SessionTimeout)
			if time.Since(session.LastActive) > b.SessionTimeout {
				delete(b.UserSessions, userID)
			}
		}
		b.UserMutex.Unlock()
	}
}

// NewBot creates a new chatbot instance with the specified name and options.
// It initializes the chatbot with default values for session timeout, session cleanup,
// and concurrent access handling unless custom options are provided.
func NewBot(name string, options ...Option) *Bot {
	bot := &Bot{
		Name:             name,
		CurrentState:     "start",
		UserSessions:     make(map[string]*UserSession),
		UserMutex:        sync.Mutex{},
		FsmStates:        make(map[string]*FsmState),
		GlobalVars:       make(map[string]string),
		StateListeners:   make(map[string]ListenerFunc),
		RuleListeners:    make(map[string]ListenerFunc),
		SessionTimeout:   30 * time.Minute, // Default session timeout
		SessionCleanup:   1 * time.Hour,    // Default session cleanup interval
		ConcurrentAccess: false,            // Default concurrent access handling is disabled
	}

	// Apply options
	for _, option := range options {
		option(bot)
	}

	// Start session cleanup timer
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

// Example usage of adding a state to the chatbot:
//
//	bot.AddState("myState", "Welcome to my state!", transitions, rules)
func (b *Bot) AddState(name, entryMessage string, transitions []Transition, rules []Rule) {
	state := &FsmState{
		Name:         name,
		EntryMessage: entryMessage,
		Transitions:  transitions,
		Rules:        rules,
	}
	b.FsmStates[name] = state
}

// AddRuleToState adds a rule to a specific state.
func (b *Bot) AddRuleToState(stateName, name, pattern, respond string, actions []Action, errorRule *ErrorRule) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	rule := Rule{
		Name:      name,
		Pattern:   re,
		Respond:   respond,
		Actions:   actions,
		ErrorRule: errorRule,
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
// It takes the user's ID and message as input and returns a response message or an error.
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
		// Handle error, misconfigured state, or default behavior
		return "State not found", nil
	}

	// Check for error rules first at the rule level
	for _, rule := range state.Rules {
		if rule.ErrorRule != nil && rule.ErrorRule.Pattern != nil && rule.ErrorRule.Pattern.MatchString(message) {
			return rule.ErrorRule.Respond, nil
		}
	}

	for _, transition := range state.Transitions {
		if transition.Event == message {
			if transition.Target == "start" {
				session.SessionState = "start"
			} else {
				session.SessionState = transition.Target
			}
			b.CurrentState = session.SessionState
			return b.replaceVariables(b.FsmStates[b.CurrentState].EntryMessage, session.SessionVars), nil
		}
	}

	var (
		wg        sync.WaitGroup
		respChan  = make(chan string, len(state.Rules))
		errorChan = make(chan error, len(state.Rules))
	)

	for _, rule := range state.Rules {
		wg.Add(1)

		go func(rule Rule) {
			defer wg.Done()

			match := rule.Pattern.FindStringSubmatch(message)
			if match != nil {
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

				// Replace variables in the response message with session variables
				respond = b.replaceVariables(respond, session.SessionVars)

				// Call state listener if available
				if listener, ok := b.StateListeners[state.Name]; ok {
					listener(userID, message, session, b)
				}

				// Call rule listener if available
				if listener, ok := b.RuleListeners[rule.Name]; ok {
					listener(userID, message, session, b)
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

	// Default response when no transitions or rules match
	return b.replaceVariables(state.EntryMessage, session.SessionVars), nil
}

// replaceVariables replaces variables in the text with their session values and global variables.
func (b *Bot) replaceVariables(text string, vars VariableMap) string {
	// Replace variables in the text with session values
	for name, value := range vars {
		placeholder := fmt.Sprintf("{{%s}}", name)
		text = strings.ReplaceAll(text, placeholder, value)
	}

	// Replace bot variables with global values
	for name, value := range b.GlobalVars {
		placeholder := fmt.Sprintf("{{bot.%s}}", name)
		text = strings.ReplaceAll(text, placeholder, value)
	}

	return text
}

// MoveToState allows moving the user to a specified state.
func (b *Bot) MoveToState(userID, newState string) {
	session, ok := b.UserSessions[userID]
	if !ok {
		session = &UserSession{
			SessionVars:  make(VariableMap),
			SessionState: b.CurrentState,
		}
		b.UserSessions[userID] = session
	}

	session.SessionState = newState
	b.CurrentState = newState
}
