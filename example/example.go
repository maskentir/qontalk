// Package main is an example program demonstrating the use of the qontalk package
// and finite state machine (fsm) package for creating and managing a Finite State Machine
// and interacting with the Qontak SDK.
package main

import (
	"fmt"
	"regexp"

	"github.com/maskentir/qontalk/fsm"
	"github.com/maskentir/qontalk/qontak"
)

// main is the entry point of the example program.
func main() {
	exampleFSM()
	exampleQontak()
}

// exampleFSM demonstrates how to create and use a Finite State Machine (FSM) using the fsm package.
func exampleFSM() {
	bot := fsm.NewBot("ChatBot")

	bot.AddState("start", "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", []fsm.Transition{
		{Event: "1", Target: "view_growth_history"},
		{Event: "2", Target: "update_growth_data"},
	}, []fsm.Rule{}, fsm.Rule{})

	bot.AddState("view_growth_history", "Growth history of your child: Name: {{child_name}} Height: {{height}} Weight: {{weight}} Month: {{month}}", []fsm.Transition{
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{}, fsm.Rule{
		Name:    "custom_error",
		Pattern: regexp.MustCompile("error"),
		Respond: "Custom error message for view_growth_history state.",
	})

	bot.AddState("update_growth_data", "Please provide the growth information for your child. Use this template e.g., 'Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm'", []fsm.Transition{
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{}, fsm.Rule{
		Name:    "custom_error",
		Pattern: regexp.MustCompile("error"),
		Respond: "Custom error message for update_growth_data state.",
	})

	bot.AddRuleToState("update_growth_data", "rule_update_growth_data", `Month: (?P<month>.+) Child's name: (?P<child_name>.+) Weight: (?P<weight>.+) kg Height: (?P<height>.+) cm`, "Thank you for updating {{child_name}}'s growth in {{month}} with height {{height}} and weight {{weight}}", nil)

	messages := []string{
		"2",
		"Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm",
		"error",
	}

	for _, message := range messages {
		response, err := bot.ProcessMessage("user1", message)
		if err != nil {
			fmt.Printf("Error processing message '%s': %v\n", message, err)
		} else {
			fmt.Printf("User1: %s\n", message)
			fmt.Printf("Bot: %s\n", response)
		}
	}
}

// exampleQontak demonstrates how to interact with the Qontak SDK using the qontak package.
func exampleQontak() {
	// Create QontakSDK instance
	sdkBuilder := qontak.NewQontakSDKBuilder().Build()

	// Authenticate with credentials
	err := sdkBuilder.Authenticate()
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return
	}

	// Create message interactions builder
	interactionsBuilder := qontak.NewSendMessageInteractionsBuilder().
		WithReceiveMessageFromAgent(true).
		WithStatusMessage(true).
		WithURL("https://example.com").
		Build()

	// Send message interactions
	err = sdkBuilder.SendMessageInteractions(interactionsBuilder)
	if err != nil {
		fmt.Println("Failed to send interactions:", err)
	}

	// Create interactive message builder
	interactiveBuilder := qontak.NewSendInteractiveMessageBuilder().
		WithRoomID("room123").
		WithInteractiveData(qontak.InteractiveData{
			Body: "Hello, World!",
			Buttons: []qontak.Button{
				{ID: "btn1", Title: "Button 1"},
				{ID: "btn2", Title: "Button 2"},
			},
		}).
		Build()

	// Send interactive message
	err = sdkBuilder.SendInteractiveMessage(interactiveBuilder)
	if err != nil {
		fmt.Println("Failed to send interactive message:", err)
	}

	// Create WhatsApp message builder
	whatsappMessageBuilder := qontak.NewWhatsAppMessageBuilder().
		WithRoomID("room123").
		WithMessage("Hello, this is a message!").
		Build()

	// Send WhatsApp message
	err = sdkBuilder.SendWhatsAppMessage(whatsappMessageBuilder)
	if err != nil {
		fmt.Println("Failed to send WhatsApp message:", err)
	}

	// Create Direct WhatsApp Broadcast builder
	directWhatsAppBroadcastBuilder := qontak.NewDirectWhatsAppBroadcastBuilder().
		WithToName("John Doe").
		WithToNumber("123456789").
		WithMessageTemplateID("template123").
		WithChannelIntegrationID("integration456").
		WithLanguage("en").
		AddHeaderParam("url", "https://example.com/sample.pdf").
		AddHeaderParam("filename", "sample.pdf").
		AddBodyParam("1", "Lorem Ipsum", "customer_name").
		AddButton(qontak.ButtonMessage{Index: "0", Type: "url", Value: "paymentUniqNumber"}).
		Build()

	// Send Direct WhatsApp Broadcast
	err = sdkBuilder.SendDirectWhatsAppBroadcast(directWhatsAppBroadcastBuilder)
	if err != nil {
		fmt.Println("Failed to send Direct WhatsApp Broadcast:", err)
	}

	// Get WhatsApp Templates
	templates, err := sdkBuilder.GetWhatsAppTemplates()
	if err != nil {
		fmt.Println("Failed to get WhatsApp Templates:", err)
	} else {
		fmt.Println("WhatsApp Templates:", templates)
	}
}
