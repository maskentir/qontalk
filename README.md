# :unicorn: Qontalk: A Go SDK for Qontak

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go version](https://img.shields.io/badge/go-1.16%2B-blue.svg)

Qontalk is a Go SDK that provides provides a unified Go SDK for seamless interaction with both the Qontak and FSM APIs, allowing you to interact with the Qontak platform. This SDK includes functionalities to authenticate, send message interactions, send interactive messages, send WhatsApp messages, send Direct WhatsApp Broadcasts, and retrieve WhatsApp Templates.

For more detailed information on using the Qontak Go SDK, check out the [GODOC.md](GODOC.md) file.

### Overview :bulb:

The qontalk package combines the functionality of the Qontak and FSM packages to provide a single, comprehensive SDK for both Qontak messaging and Finite State Machine \(FSM\) operations. This package allows you to effortlessly work with Qontak for messaging while simultaneously building, managing, and executing FSMs within your applications.

### Qontak Integration :cactus:

The Qontak integration within qontalk enables you to:

- Authenticate with Qontak's services using various authentication methods.
- Send messages to customers and agents, including text, images, and interactive messages.
- Manage WhatsApp templates and send WhatsApp messages.
- Perform Direct WhatsApp Broadcasts with customization options.

You can utilize these features to enhance your messaging capabilities and communication with customers and agents through Qontak's platform.

### Finite State Machine \(FSM\) Integration :rocket:

The FSM integration allows you to create, manage, and execute Finite State Machines within your application. You can define custom states, events, transitions, and callbacks to control the flow of your application based on specific conditions.

Key features of the FSM integration include:

- Creating FSM instances with custom initial states and transitions.
- Sending events to trigger state transitions.
- Managing FSM lifecycle, including starting and stopping FSM execution.
- Defining global callbacks to handle state transitions and events.

This FSM integration empowers you to build complex, stateful applications with ease.

## Installation

To install the Qontalk SDK, you can use `go get`:

```sh
go get github.com/maskentir/qontalk
```

## Documentation

You can find the full documentation for the Qontalk SDK [here](https://pkg.go.dev/github.com/maskentir/qontalk).

## Usage

Here is a simple example of how to use the Qontalk SDK:

```go
package main

import (
    "fmt"
    "github.com/maskentir/qontalk/qontak"
    "github.com/maskentir/qontalk/fsm"
)

func main() {
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

    // Create finite state machine (FSM) with transitions
    transitions := []fsm.Transition{
        {
            From:   "StateA",
            Event:  "Event1",
            To:     "StateB",
            Action: func() error { return nil },
        },
        // Add more transitions as needed
    }

    globalCallback := func(from fsm.State, event fsm.Event, to fsm.State, params map[string]interface{}) {
        fmt.Printf("Transition from %v to %v due to event %v\n", from, to, event)
    }

    fsmInstance, err := fsm.NewFSM("StateA", transitions, globalCallback)
    if err != nil {
        fmt.Println("Failed to create FSM:", err)
        return
    }

    // Send an event to the FSM
    event := "Event1"
    params := make(map[string]interface{})
    err = fsmInstance.SendEvent(event, params)
    if err != nil {
        fmt.Println("Failed to send event to FSM:", err)
    }

    // Get the current state of the FSM
    currentState := fsmInstance.GetCurrentState()
    fmt.Println("Current state:", currentState)
}
```

## License

This library is released under the [MIT License](LICENSE).

## Contributing

Please read our [Contribution Guidelines](CONTRIBUTE.md) before submitting a pull request.

## Code of Conduct

Please follow our [Code of Conduct](CODE_OF_CONDUCT.md) when participating in this project.

## Issue and Pull Request Templates

Before creating an issue or pull request, review our [Issue Template](ISSUE_TEMPLATE.md) and [Pull Request Template](PULL_REQUEST_TEMPLATE.md).

## License

This library is released under the MIT [License](LICENSE).

## Contact

If you have any questions or feedback, please contact our support team at harunwols@gmail.com.
