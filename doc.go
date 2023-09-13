// Package qontalk provides a unified Go SDK for seamless interaction with both the Qontak and FSM APIs.
//
// # Overview
//
// The qontalk package combines the functionality of the Qontak and FSM packages
// to provide a single, comprehensive SDK for both Qontak messaging and Finite State
// Machine (FSM) operations. This package allows you to effortlessly work with Qontak
// for messaging while simultaneously building, managing, and executing FSMs within
// your applications.
//
// # Qontak Integration
//
// The Qontak integration within qontalk enables you to:
//
// - Authenticate with Qontak's services using various authentication methods.
// - Send messages to customers and agents, including text, images, and interactive messages.
// - Manage WhatsApp templates and send WhatsApp messages.
// - Perform Direct WhatsApp Broadcasts with customization options.
//
// You can utilize these features to enhance your messaging capabilities and
// communication with customers and agents through Qontak's platform.
//
// # Finite State Machine (FSM) Integration
//
// The FSM integration allows you to create, manage, and execute Finite State Machines
// within your application. You can define custom states, events, transitions, and
// callbacks to control the flow of your application based on specific conditions.
//
// Key features of the FSM integration include:
//
// - Creating FSM instances with custom initial states and transitions.
// - Sending events to trigger state transitions.
// - Managing FSM lifecycle, including starting and stopping FSM execution.
// - Defining global callbacks to handle state transitions and events.
//
// This FSM integration empowers you to build complex, stateful applications with ease.
//
// # Example Usage
//
// The following example demonstrates how to leverage the qontalk package to interact
// with both Qontak and FSM functionalities:
//
// func main() {
// 	// Create a QontalkSDK instance
// 	sdk := qontalk.NewQontalkSDKBuilder().
// 		WithClientCredentials("your-username", "your-password", "your-grant-type", "your-client-id", "your-client-secret").
// 		Build()

// 	// Authenticate with Qontak
// 	if err := sdk.Authenticate(); err != nil {
// 		fmt.Println("Authentication failed:", err)
// 		return
// 	}

// 	// Use Qontak features, send messages, etc.

// 	// Create a new chatbot instance
// 	bot := fsm.NewBot("MyChatbot")

// 	// Define states and transitions
// 	transitions := []fsm.Transition{
// 		{Event: "start", Target: "initial"},
// 		{Event: "continue", Target: "ongoing"},
// 	}

// 	bot.AddState("initial", "Welcome to the chatbot!", transitions)

// 	// Define rules and actions
// 	rulePattern := "hello"
// 	regexPattern := fmt.Sprintf("(?i)%s", regexp.QuoteMeta(rulePattern))
// 	rule := fsm.Rule{
// 		Name:    "HelloRule",
// 		Pattern: regexp.MustCompile(regexPattern),
// 		Respond: "Hello! How can I assist you?",
// 	}

// 	bot.AddRuleToState("initial", rule.Name, regexPattern, rule.Respond, nil, nil)

// 	// Process user messages
// 	response, err := bot.ProcessMessage("user123", "hello")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

//		fmt.Println("Bot Response:", response)
//	}
//
// This example showcases how you can use the qontalk package to work with Qontak
// messaging and FSM features in a single application, creating a unified experience.
//
// # Additional Resources
//
// For more detailed documentation and comprehensive usage examples, please refer to
// the individual package documentation for qontak and fsm.
package qontalk
