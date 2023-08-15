// Package qontak provides a SDK for interacting with the Qontak API.
//
// The QontakSDKBuilder type is used to build instances of the QontakSDK,
// which is a singleton for accessing the Qontak API. You can use the SDK
// to authenticate, send message interactions, and send interactive messages.
//
// Authentication
//
// The QontakSDK can be authenticated using the Authenticate method, which
// retrieves an access token for making authenticated API requests.
//
// Sending Message Interactions
//
// You can use the SendMessageInteractions method to send message interactions,
// including settings for receiving messages from agents or customers and setting
// the status message and URL.
//
// Sending Interactive Messages
//
// The SendInteractiveMessage method allows you to send interactive messages
// to a specified room ID, using interactive data.
//
// Customizing Request Strategy
//
// The QontakSDK uses a RequestStrategy interface for sending requests. The
// DefaultRequestStrategy is the default implementation of this interface, but
// you can also set a custom strategy using the SetRequestStrategy method.
//
//
// Example Usage
//
// Here's a quick example demonstrating how to authenticate and send an interactive message using the Qontak SDK:
//
//     // Create a QontakSDKBuilder with client credentials
//     builder := qontak.NewQontakSDKBuilder().WithClientCredentials(
//         "your-username", "your-password", "password", "your-client-id", "your-client-secret",
//     )
//
//     // Build the QontakSDK instance
//     sdk := builder.Build()
//
//     // Authenticate the SDK
//     err := sdk.Authenticate()
//     if err != nil {
//         fmt.Println("Authentication failed:", err)
//         return
//     }
//
//     // Send a message interaction
//     interactionBuilder := qontak.NewSendMessageInteractionsBuilder().
//         WithReceiveMessageFromAgent(true).
//         WithStatusMessage(true).
//         WithURL("https://example.com")
//     err = sdk.SendMessageInteractions(interactionBuilder.Build())
//     if err != nil {
//         fmt.Println("Failed to send message interaction:", err)
//     }
//
//     // Send an interactive message
//     interactiveBuilder := qontak.NewSendInteractiveMessageBuilder().
//         WithRoomID("room123").
//         WithInteractiveData(interactiveData)
//     err = sdk.SendInteractiveMessage(interactiveBuilder.Build())
//     if err != nil {
//         fmt.Println("Failed to send interactive message:", err)
//     }
//
package qontak
