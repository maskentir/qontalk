// Package qontak provides a Go SDK for interacting with the Qontak API.
//
// Documentation
//
// This package offers a Go SDK for accessing the Qontak API. It provides a
// convenient way to authenticate, send message interactions, and send
// interactive messages using the Qontak SDK.
//
// Overview
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
// Examples
//
// The following example demonstrates how to use the SDK to send a message
// interaction and an interactive message:
//
//   // Create QontakSDK instance
//   sdkBuilder := NewQontakSDKBuilder().Build()
//
//   // Authenticate with credentials
//   err := sdkBuilder.Authenticate()
//   if err != nil {
//       fmt.Println("Authentication failed:", err)
//       return
//   }
//
//   // Create message interactions builder
//   interactionsBuilder := NewSendMessageInteractionsBuilder().
//       WithReceiveMessageFromAgent(true).
//       WithReceiveMessageFromCustomer(true).
//       WithStatusMessage(true).
//       WithURL("https://example.com").
//       Build()
//
//   // Send message interactions
//   err = sdkBuilder.SendMessageInteractions(interactionsBuilder)
//   if err != nil {
//       fmt.Println("Failed to send interactions:", err)
//   }
//
//   // Create interactive message builder
//   interactiveBuilder := NewSendInteractiveMessageBuilder().
//       WithRoomID("room123").
//       WithInteractiveData(InteractiveData{
//           Body: "Hello, World!",
//           Buttons: []Button{
//               {ID: "btn1", Title: "Button 1"},
//               {ID: "btn2", Title: "Button 2"},
//           },
//       }).
//       Build()
//
//   // Send interactive message
//   err = sdkBuilder.SendInteractiveMessage(interactiveBuilder)
//   if err != nil {
//       fmt.Println("Failed to send interactive message:", err)
//   }
package qontak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// QontakSDKBuilder is a builder to create QontakSDK
type QontakSDKBuilder struct {
	username     string
	password     string
	grantType    string
	clientID     string
	clientSecret string
}

// NewQontakSDKBuilder creates a new instance of QontakSDKBuilder
func NewQontakSDKBuilder() *QontakSDKBuilder {
	return &QontakSDKBuilder{}
}

// WithClientCredentials sets client credentials for the builder.
// Example:
// builder.WithClientCredentials("your-username", "your-password", "password", "your-client-id", "your-client-secret")
func (b *QontakSDKBuilder) WithClientCredentials(
	username, password, grantType, clientID, clientSecret string,
) *QontakSDKBuilder {
	b.username = username
	b.password = password
	b.grantType = grantType
	b.clientID = clientID
	b.clientSecret = clientSecret
	return b
}

// Build builds QontakSDK from the builder.
// Example:
// sdk := builder.Build()
func (b *QontakSDKBuilder) Build() *QontakSDK {
	return &QontakSDK{
		BaseURL:         "https://service-chat.qontak.com/api/open/v1",
		Username:        b.username,
		Password:        b.password,
		GrantType:       b.grantType,
		ClientID:        b.clientID,
		ClientSecret:    b.clientSecret,
		RequestStrategy: &DefaultRequestStrategy{},
	}
}

// QontakSDK is a singleton for accessing Qontak API
type QontakSDK struct {
	BaseURL         string
	Username        string
	Password        string
	GrantType       string
	ClientID        string
	ClientSecret    string
	AccessToken     string
	RequestStrategy RequestStrategy
}

// Authenticate authenticates the SDK with the provided credentials.
// Example:
// err := sdk.Authenticate()
func (sdk *QontakSDK) Authenticate() error {
	authURL := fmt.Sprintf("%s/oauth/token", sdk.BaseURL)

	data := map[string]interface{}{
		"username":      sdk.Username,
		"password":      sdk.Password,
		"grant_type":    sdk.GrantType,
		"client_id":     sdk.ClientID,
		"client_secret": sdk.ClientSecret,
	}

	resp, err := sdk.RequestStrategy.Post(authURL, data)
	if err != nil {
		return err
	}

	accessToken, ok := resp["access_token"].(string)
	if !ok {
		return fmt.Errorf("authentication failed")
	}

	sdk.AccessToken = accessToken
	return nil
}

// SendMessageInteractions sends message interactions.
// Example:
// builder := NewSendMessageInteractionsBuilder().WithReceiveMessageFromAgent(true).WithStatusMessage(true).WithURL("https://example.com")
// err := sdk.SendMessageInteractions(builder.Build())
func (sdk *QontakSDK) SendMessageInteractions(builder SendMessageInteractions) error {
	interactionURL := fmt.Sprintf("%s/message_interactions", sdk.BaseURL)

	data := map[string]interface{}{
		"receive_message_from_agent":    builder.ReceiveMessageFromAgent,
		"receive_message_from_customer": builder.ReceiveMessageFromCustomer,
		"status_message":                builder.StatusMessage,
		"url":                           builder.URL,
	}

	_, err := sdk.RequestStrategy.Put(interactionURL, data)
	return err
}

// SendInteractiveMessage sends an interactive message.
// Example:
// builder := NewSendInteractiveMessageBuilder().WithRoomID("room123").WithInteractiveData(interactiveData)
// err := sdk.SendInteractiveMessage(builder.Build())
func (sdk *QontakSDK) SendInteractiveMessage(builder SendInteractiveMessage) error {
	url := fmt.Sprintf("%s/messages/whatsapp/interactive_message/bot", sdk.BaseURL)

	data := map[string]interface{}{
		"room_id":     builder.RoomID,
		"type":        builder.Type,
		"interactive": builder.Interactive,
	}

	_, err := sdk.RequestStrategy.Post(url, data)
	return err
}

// RequestStrategy is a strategy interface for sending requests
type RequestStrategy interface {
	// Post sends a POST request with the default strategy.
	// Example:
	// resp, err := drs.Post(url, data)
	Post(url string, data map[string]interface{}) (map[string]interface{}, error)
	// Put sends a PUT request with the default strategy.
	// Example:
	// resp, err := drs.Put(url, data)
	Put(url string, data map[string]interface{}) (map[string]interface{}, error)
}

// DefaultRequestStrategy is the default implementation of RequestStrategy
type DefaultRequestStrategy struct {
	AccessToken string
}

// Post sends a POST request with the default strategy.
// Example:
// resp, err := drs.Post(url, data)
func (drs *DefaultRequestStrategy) Post(
	url string,
	data map[string]interface{},
) (map[string]interface{}, error) {
	payloadBytes, _ := json.Marshal(data)
	payload := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", drs.AccessToken))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// Put sends a PUT request with the default strategy.
// Example:
// resp, err := drs.Put(url, data)
func (drs *DefaultRequestStrategy) Put(
	url string,
	data map[string]interface{},
) (map[string]interface{}, error) {
	payloadBytes, _ := json.Marshal(data)
	payload := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if drs.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+drs.AccessToken)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody, nil
}

// SetRequestStrategy sets the request strategy in QontakSDK.
// Example:
// sdk.SetRequestStrategy(&CustomRequestStrategy{})
func (sdk *QontakSDK) SetRequestStrategy(strategy RequestStrategy) {
	sdk.RequestStrategy = strategy
}
