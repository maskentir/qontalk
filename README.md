# :unicorn: Qontalk: A Go SDK for Qontak

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go version](https://img.shields.io/badge/go-1.16%2B-blue.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/maskentir/qontalk.svg)](https://pkg.go.dev/github.com/maskentir/qontalk)

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
    "github.com/maskentir/qontalk"
    "github.com/maskentir/qontalk/fsm"
)

func main() {
    // Create a QontalkSDK instance
    sdk := qontalk.NewQontalkSDKBuilder().
        WithClientCredentials("your-username", "your-password", "your-grant-type", "your-client-id", "your-client-secret").
        Build()

    // Authenticate with Qontak
    if err := sdk.Authenticate(); err != nil {
        fmt.Println("Authentication failed:", err)
        return
    }

    // Use Qontak features, send messages, etc.

    // Create an FSM instance
    fsm := fsm.NewBot("ChatBot")

    fsm.AddState("start", "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", []fsm.Transition{
        {Event: "1", Target: "view_growth_history"},
        {Event: "2", Target: "update_growth_data"},
    }, []fsm.Rule{}, fsm.Rule{})

    fsm.AddState("view_growth_history", "Growth history of your child: Name: {{child_name}} Height: {{height}} Weight: {{weight}} Month: {{month}}", []fsm.Transition{
        {Event: "exit", Target: "start"},
    }, []fsm.Rule{}, fsm.Rule{
        Name:    "custom_error",
        Pattern: regexp.MustCompile("error"),
        Respond: "Custom error message for view_growth_history state.",
    })

    fsm.AddState("update_growth_data", "Please provide the growth information for your child. Use this template e.g., 'Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm'", []fsm.Transition{
        {Event: "exit", Target: "start"},
    }, []fsm.Rule{}, fsm.Rule{
        Name:    "custom_error",
        Pattern: regexp.MustCompile("error"),
        Respond: "Custom error message for update_growth_data state.",
    })

    fsm.AddRuleToState("update_growth_data", "rule_update_growth_data", `Month: (?P<month>.+) Child's name: (?P<child_name>.+) Weight: (?P<weight>.+) kg Height: (?P<height>.+) cm`, "Thank you for updating {{child_name}}'s growth in {{month}} with height {{height}} and weight {{weight}}", nil)

    messages := []string{
        "2",
        "Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm",
        "error",
    }

    for _, message := range messages {
        response, err := fsm.ProcessMessage("user1", message)
        if err != nil {
            fmt.Printf("Error processing message '%s': %v\n", message, err)
        } else {
            fmt.Printf("User1: %s\n", message)
            fmt.Printf("Bot: %s\n", response)
        }
    }
}
```

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
