package fsm_test

import (
	"regexp"
	"testing"

	"github.com/maskentir/qontalk/fsm"
)

func TestBot_ProcessMessage(t *testing.T) {
	bot := fsm.NewBot("TestBot")

	bot.AddState("start", "Hi Ayah Bunda! Balas dengan menu di bawah sesuai kebutuhan Ayah Bunda:\n1 Riwayat tumbuh kembang anak\n2 Update tumbuh kembang anak\nContoh: ketik '1' jika Anda ingin mengetahui data pertumbuhan anak Anda", []fsm.Transition{
		{Event: "1", Target: "view_growth_history"},
		{Event: "2", Target: "update_growth_data"},
	}, []fsm.Rule{}, fsm.Rule{})

	bot.AddState("view_growth_history", "Riwayat tumbuh & kembang anak Ayah Bunda: Nama: {{child_name}} TB: {{tb}} BB: {{bb}} Bulan: {{month}}", []fsm.Transition{
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{}, fsm.Rule{
		Name:    "custom_error",
		Pattern: regexp.MustCompile("error"),
		Respond: "Custom error message for view_growth_history state.",
	})

	bot.AddState("update_growth_data", "Silahkan Ayah Bunda memberikan informasi pertumbuhan anak Anda. Ikuti template ini e.g; 'Bulan: Januari Nama anak: Harun Nur Rasyid BB: 30,5 kg TB: 89,1 cm'", []fsm.Transition{
		{Event: "exit", Target: "start"},
	}, []fsm.Rule{}, fsm.Rule{
		Name:    "custom_error",
		Pattern: regexp.MustCompile("error"),
		Respond: "Custom error message for update_growth_data state.",
	})

	bot.AddRuleToState("update_growth_data", "rule_update_growth_data", `Bulan: (?P<month>.+) Nama anak: (?P<child_name>.+) BB: (?P<bb>.+) kg TB: (?P<tb>.+) cm`, "Terimakasih saya sudah mengupdate pertumbuhan {{child_name}} di bulan {{month}} dengan TB {{tb}} cm dan BB {{bb}} kg", nil)

	testCases := []struct {
		UserID       string
		Message      string
		ExpectedResp string
		ExpectedErr  error
	}{
		{
			UserID:       "user1",
			Message:      "2",
			ExpectedResp: "Silahkan Ayah Bunda memberikan informasi pertumbuhan anak Anda. Ikuti template ini e.g; 'Bulan: Januari Nama anak: Harun Nur Rasyid BB: 30,5 kg TB: 89,1 cm'",
			ExpectedErr:  nil,
		},
		{
			UserID:       "user1",
			Message:      "Bulan: Januari Nama anak: Waode Melawati BB: 33,5 kg TB: 69,1 cm",
			ExpectedResp: "Terimakasih saya sudah mengupdate pertumbuhan Waode Melawati di bulan Januari dengan TB 69,1 cm dan BB 33,5 kg",
			ExpectedErr:  nil,
		},
		{
			UserID:       "user1",
			Message:      "error",
			ExpectedResp: "Custom error message for update_growth_data state.",
			ExpectedErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.UserID+"_"+tc.Message, func(t *testing.T) {
			resp, err := bot.ProcessMessage(tc.UserID, tc.Message)
			if err != tc.ExpectedErr {
				t.Errorf("Expected error %v, got %v", tc.ExpectedErr, err)
			}
			if resp != tc.ExpectedResp {
				t.Errorf("Expected response '%s', got '%s'", tc.ExpectedResp, resp)
			}
		})
	}
}
