package telegram

import (
	"os"
	"testing"
)

func TestTodayIsReportCheckDay(t *testing.T) {
	os.Setenv("DAYS_FOR_REPORT_CHECK", "Saturday,Sunday,Monday")

	for _, dayOk := range []string{"Sunday", "Monday", "Saturday"} {
		t.Run(dayOk+" day", func(t *testing.T) {
			if !todayIsReportCheckDay(dayOk) {
				t.Errorf("todayIsReportCheckDay for value %s must return true", dayOk)
			}
		})
	}

	for _, day := range []string{"Tuesday", "Wednesday", "Thursday", "Friday"} {
		t.Run(day+" day", func(t *testing.T) {
			if todayIsReportCheckDay(day) {
				t.Errorf("todayIsReportCheckDay for value %s must return false", day)
			}
		})
	}
}

func TestGetChatId(t *testing.T) {
	os.Setenv("BOT_CHAT_ID", "543210")
	chatID := getChatID()

	if chatID != 543210 {
		t.Errorf("The result of getChatID function must be 543210, but %v returned", chatID)
	}
}
