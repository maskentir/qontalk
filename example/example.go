package main

import (
	"fmt"

	"github.com/maskentir/qontalk/fsm"
	"github.com/maskentir/qontalk/qontak"
)

func main() {
	exampleFSM()
	exampleQontak()
}

func exampleFSM() {
	// Define custom state types.
	type MyState int
	const (
		StateA MyState = iota
		StateB
		StateC
	)

	// Define custom event types.
	type MyEvent int
	const (
		EventX MyEvent = iota
		EventY
	)

	// Define a callback function to execute when transitioning.
	callback := func(from fsm.State, event fsm.Event, to fsm.State, params map[string]interface{}) {
		fmt.Printf("Transition from %v to %v due to event %v\n", from, to, event)
	}

	// Create an FSM instance with an initial state, transitions, and the callback.
	transitions := []fsm.Transition{
		{From: StateA, Event: EventX, To: StateB},
		{From: StateB, Event: EventY, To: StateC},
	}

	fsmInstance, err := fsm.NewFSM(StateA, transitions, callback)
	if err != nil {
		fmt.Println("Error creating FSM:", err)
		return
	}

	// Send events to trigger transitions.
	fsmInstance.SendEvent(EventX, nil)
	fsmInstance.SendEvent(EventY, nil)

	// Get the current state.
	currentState := fsmInstance.GetCurrentState()
	fmt.Println("Current State:", currentState)

	// Stop the FSM (this will wait for all goroutines to complete).
	fsmInstance.Stop()

	fmt.Println("FSM Stopped")
}

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
