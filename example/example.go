// Package main is an example of using the QontakSDK to interact with the Qontak platform.
// It demonstrates how to create an instance of the QontakSDK, authenticate it, and send
// message interactions as well as interactive messages.
package main

import (
	"fmt"
	"log"

	qontak "github.com/maskentir/qontalk/qontak"
)

// ExampleMain demonstrates how to use the QontakSDK to interact with the Qontak platform.
// Create an instance of QontakSDK using the builder pattern
// sdk := qontak.NewQontakSDKBuilder().
// 	WithClientCredentials("your-username", "your-password", "your-grant-type", "your-client-id", "your-client-secret").
// 	Build()

// // Authenticate the SDK using the provided credentials
// if err := sdk.Authenticate(); err != nil {
// 	log.Fatal("Authentication failed:", err)
// }

//  Send message interactions configuration
// interactionsBuilder := qontak.SendMessageInteractions{
// 	ReceiveMessageFromAgent:    true,
// 	ReceiveMessageFromCustomer: true,
// 	StatusMessage:              true,
// 	URL:                        "https://example.com",
// }
// Send the configured message interactions to the Qontak platform
// if err := sdk.SendMessageInteractions(interactionsBuilder); err != nil {
// 	log.Println("Sending message interactions failed:", err)
// } else {
// 	fmt.Println("Message interactions sent successfully")
// }

// Send an interactive message configuration
// interactiveBuilder := qontak.SendInteractiveMessage{
// 	RoomID: "room123",
// 	Type:   "type123",
// 	Interactive: qontak.InteractiveData{
// 		Header: &qontak.InteractiveHeader{
// 			Format:   "json",
// 			Text:     "Header Text",
// 			Link:     "https://example.com",
// 			Filename: "file.txt",
// 		},
// 		Body: "Body Text",
// 		Buttons: []qontak.Button{
// 			{ID: "btn1", Title: "Button 1"},
// 			{ID: "btn2", Title: "Button 2"},
// 		},
// 	},
// }
//
// Send the configured interactive message to the Qontak platform
// if err := sdk.SendInteractiveMessage(interactiveBuilder); err != nil {
// 	log.Println("Sending interactive message failed:", err)
// } else {
// 	fmt.Println("Interactive message sent successfully")
// }
func ExampleMain() {
	// Create an instance of QontakSDK using the builder pattern
	sdk := qontak.NewQontakSDKBuilder().
		WithClientCredentials("your-username", "your-password", "your-grant-type", "your-client-id", "your-client-secret").
		Build()

	// Authenticate the SDK using the provided credentials
	if err := sdk.Authenticate(); err != nil {
		log.Fatal("Authentication failed:", err)
	}

	// Send message interactions configuration
	interactionsBuilder := qontak.SendMessageInteractions{
		ReceiveMessageFromAgent:    true,
		ReceiveMessageFromCustomer: true,
		StatusMessage:              true,
		URL:                        "https://example.com",
	}
	// Send the configured message interactions to the Qontak platform
	if err := sdk.SendMessageInteractions(interactionsBuilder); err != nil {
		log.Println("Sending message interactions failed:", err)
	} else {
		fmt.Println("Message interactions sent successfully")
	}

	// Send an interactive message configuration
	interactiveBuilder := qontak.SendInteractiveMessage{
		RoomID: "room123",
		Type:   "type123",
		Interactive: qontak.InteractiveData{
			Header: &qontak.InteractiveHeader{
				Format:   "json",
				Text:     "Header Text",
				Link:     "https://example.com",
				Filename: "file.txt",
			},
			Body: "Body Text",
			Buttons: []qontak.Button{
				{ID: "btn1", Title: "Button 1"},
				{ID: "btn2", Title: "Button 2"},
			},
		},
	}
	// Send the configured interactive message to the Qontak platform
	if err := sdk.SendInteractiveMessage(interactiveBuilder); err != nil {
		log.Println("Sending interactive message failed:", err)
	} else {
		fmt.Println("Interactive message sent successfully")
	}
}

func main() {
	// Call the ExampleMain function to demonstrate the usage of the QontakSDK.
	ExampleMain()
}
