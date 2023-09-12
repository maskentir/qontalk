package fsm_test

import (
	"testing"
	"time"

	"github.com/maskentir/qontalk/fsm"
)

func TestProcessMessage(t *testing.T) {
	bot := fsm.NewBot("TestBot")

	bot.AddState("start", "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", []fsm.Transition{
		{Event: "1", Target: "view_growth_history"},
		{Event: "2", Target: "update_growth_data"},
	}, []fsm.Rule{})

	bot.AddState("view_growth_history", "Growth history of your child: Name: {{child_name}} Height: {{height}} Weight: {{weight}} Month: {{month}}", []fsm.Transition{
		{Event: "1", Target: "view_growth_history"},
		{Event: "2", Target: "update_growth_data"},
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{})

	bot.AddState("update_growth_data", "Please provide the growth information for your child. Use this template e.g., 'Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm'", []fsm.Transition{
		{Event: "1", Target: "view_growth_history"},
		{Event: "2", Target: "update_growth_data"},
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{})

	bot.AddRuleToState("update_growth_data", "rule_update_growth_data", `Month: (?P<month>.+) Child's name: (?P<child_name>.+) Weight: (?P<weight>.+) kg Height: (?P<height>.+) cm`, "Thank you for updating {{child_name}}'s growth in {{month}} with height {{height}} cm and weight {{weight}} kg", nil, nil)

	tests := []struct {
		UserID      string
		Message     string
		Expected    string
		ExpectError bool
	}{
		{UserID: "user1", Message: "1", Expected: "Growth history of your child: Name: {{child_name}} Height: {{height}} Weight: {{weight}} Month: {{month}}", ExpectError: false},

		{UserID: "user1", Message: "2", Expected: "Please provide the growth information for your child. Use this template e.g., 'Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm'", ExpectError: false},

		{UserID: "user1", Message: "Month: January Child's name: John Weight: 30.5 kg Height: 89.1 cm", Expected: "Thank you for updating John's growth in January with height 89.1 cm and weight 30.5 kg", ExpectError: false},
	}
	for _, test := range tests {
		response, err := bot.ProcessMessage(test.UserID, test.Message)

		if (err != nil) != test.ExpectError {
			t.Errorf("Expected error: %v, but got error: %v", test.ExpectError, err)
		}

		if response != test.Expected {
			t.Errorf("For User: %s, Message: %s - Expected: %s, but got: %s", test.UserID, test.Message, test.Expected, response)
		}
	}
}

func TestAdvancedFeatures(t *testing.T) {
	bot := fsm.NewBot("TestBot", fsm.WithSessionCleanup(1*time.Second), fsm.WithSessionTimeout(2*time.Second)) // Session cleanup setiap 1 detik untuk pengujian

	bot.AddState("start", "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", []fsm.Transition{
		{Event: "custom", Target: "custom_state"},
	}, []fsm.Rule{})

	bot.AddState("custom_state", "This is a custom state.", []fsm.Transition{
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{})

	bot.AddRuleToState("custom_state", "rule_custom", `custom pattern`, "Custom response", nil, nil)

	bot.AddListenerToState("custom_state", func(userID string, message string, session *fsm.UserSession, bot *fsm.Bot) {
	})

	bot.AddListenerToRule("rule_custom", func(userID string, message string, session *fsm.UserSession, bot *fsm.Bot) {
	})

	tests := []struct {
		UserID      string
		Message     string
		Expected    string
		ExpectError bool
	}{
		{UserID: "user1", Message: "custom", Expected: "This is a custom state.", ExpectError: false},

		{UserID: "user1", Message: "exit", Expected: "Hi there! Reply with one of the following options:\n1 View growth history\n2 Update growth data\nExample: type '1' if you want to view your child's growth history.", ExpectError: false},

		{UserID: "user1", Message: "custom", Expected: "This is a custom state.", ExpectError: false},

		{UserID: "user1", Message: "custom pattern", Expected: "Custom response", ExpectError: false},

		{UserID: "user1", Message: "", Expected: "This is a custom state.", ExpectError: false},
	}
	for _, test := range tests {
		response, err := bot.ProcessMessage(test.UserID, test.Message)

		if (err != nil) != test.ExpectError {
			t.Errorf("Expected error: %v, but got error: %v", test.ExpectError, err)
		}

		if response != test.Expected {
			t.Errorf("For User: %s, Message: %s - Expected: %s, but got: %s", test.UserID, test.Message, test.Expected, response)
		}
	}

	time.Sleep(5 * time.Second)

	bot.UserMutex.Lock()
	defer bot.UserMutex.Unlock()
	_, sessionExists := bot.UserSessions["user1"]
	if sessionExists {
		t.Errorf("Expected session 'user1' to be deleted after expiration, but it still exists")
	}
}
