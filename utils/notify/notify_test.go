package notify

import (
	"os"
	"testing"
	"time"
)

// TestSendEditDelete tests the Send, Edit, and Delete functions for both Discord and Telegram.
func TestSendEditDelete(t *testing.T) {
	discordID, ok := os.LookupEnv("DISCORD_ID")
	if !ok {
		t.Error("Discord webhook ID must be set in the environment variable 'DISCORD_ID'")
	}
	discordToken, ok := os.LookupEnv("DISCORD_TOKEN")
	if !ok {
		t.Error("Discord webhook token must be set in the environment variable 'DISCORD_TOKEN'")
	}

	telegramID, ok := os.LookupEnv("TELEGRAM_ID")
	if !ok {
		t.Error("Telegram chat ID must be set in the environment variable 'TELEGRAM_ID'")
	}
	telegramToken, ok := os.LookupEnv("TELEGRAM_TOKEN")
	if !ok {
		t.Error("Telegram bot token must be set in the environment variable 'TELEGRAM_TOKEN'")
	}

	msg := "This is a **bold**, _italic_, [inline URL](http://www.example.com/), `code`."

	discordClient := NewDiscord(discordID, discordToken)
	telegramClient := NewTelegram(telegramToken, telegramID)

	// Send messages
	dcID, err := discordClient.Send(msg)
	if err != nil {
		t.Errorf("Failed to send Discord message: %v", err)
	}
	tgID, err := telegramClient.Send(msg)
	if err != nil {
		t.Errorf("Failed to send Telegram message: %v", err)
	}

	// Wait for messages to be sent
	time.Sleep(3 * time.Second)

	// Edit messages
	if err := discordClient.Edit(dcID, "Redacted Discord Message"); err != nil {
		t.Errorf("Failed to edit Discord message: %v", err)
	}
	if err := telegramClient.Edit(tgID, "Redacted Telegram Message"); err != nil {
		t.Errorf("Failed to edit Telegram message: %v", err)
	}

	// Wait for messages to be edited
	time.Sleep(3 * time.Second)

	// Delete messages
	if err := discordClient.Delete(dcID); err != nil {
		t.Errorf("Failed to delete Discord message: %v", err)
	}
	if err := telegramClient.Delete(tgID); err != nil {
		t.Errorf("Failed to delete Telegram message: %v", err)
	}
}
