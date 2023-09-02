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
//	package main
//
//	import (
//	    "fmt"
//	    "github.com/maskentir/qontalk"
//	    "github.com/maskentir/qontalk/fsm"
//	)
//
//	func main() {
//	    // Create a QontalkSDK instance
//	    sdk := qontalk.NewQontalkSDKBuilder().
//	        WithClientCredentials("your-username", "your-password", "your-grant-type", "your-client-id", "your-client-secret").
//	        Build()
//
//	    // Authenticate with Qontak
//	    if err := sdk.Authenticate(); err != nil {
//	        fmt.Println("Authentication failed:", err)
//	        return
//	    }
//
//	    // Use Qontak features, send messages, etc.
//
//	    // Create an FSM instance
//	    fsm := fsm.NewBot("ChatBot")
//
//	    fsm.AddState("start", "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", []fsm.Transition{
//	        {Event: "1", Target: "view_growth_history"},
//	        {Event: "2", Target: "update_growth_data"},
//	    }, []fsm.Rule{}, fsm.Rule{})
//
//	    fsm.AddState("view_growth_history", "Growth history of your child: Name: {{child_name}} Height: {{height}} Weight: {{weight}} Month: {{month}}", []fsm.Transition{
//	        {Event: "exit", Target: "start"},
//	    }, []fsm.Rule{}, fsm.Rule{
//	        Name:    "custom_error",
//	        Pattern: regexp.MustCompile("error"),
//	        Respond: "Custom error message for view_growth_history state.",
//	    })
//
//	    fsm.AddState("update_growth_data", "Please provide the growth information for your child. Use this template e.g., 'Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm'", []fsm.Transition{
//	        {Event: "exit", Target: "start"},
//	    }, []fsm.Rule{}, fsm.Rule{
//	        Name:    "custom_error",
//	        Pattern: regexp.MustCompile("error"),
//	        Respond: "Custom error message for update_growth_data state.",
//	    })
//
//	    fsm.AddRuleToState("update_growth_data", "rule_update_growth_data", `Month: (?P<month>.+) Child's name: (?P<child_name>.+) Weight: (?P<weight>.+) kg Height: (?P<height>.+) cm`, "Thank you for updating {{child_name}}'s growth in {{month}} with height {{height}} and weight {{weight}}", nil)
//
//	    messages := []string{
//	        "2",
//	        "Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm",
//	        "error",
//	    }
//
//	    for _, message := range messages {
//	        response, err := fsm.ProcessMessage("user1", message)
//	        if err != nil {
//	            fmt.Printf("Error processing message '%s': %v\n", message, err)
//	        } else {
//	            fmt.Printf("User1: %s\n", message)
//	            fmt.Printf("Bot: %s\n", response)
//	        }
//	    }
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
