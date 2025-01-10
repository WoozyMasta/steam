package notify

import (
	"bytes"
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

// Telegram represents a client for interacting with the Telegram Bot API.
type Telegram struct {
	botToken string
	chatID   string
}

// NewTelegram creates a new instance of a Telegram client.
//
// Parameters:
//   - botToken: The Telegram bot token.
//   - chatID: The chat ID where messages will be sent.
//
// Returns:
//   - A pointer to a Telegram instance.
func NewTelegram(botToken, chatID string) *Telegram {
	return &Telegram{
		botToken: botToken,
		chatID:   chatID,
	}
}

// tgPayload represents the payload structure for Telegram API requests.
type tgPayload struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
	MessageID uint64 `json:"message_id,omitempty"`
}

// Send sends a message to a Telegram chat.
//
// Parameters:
//   - msg: The message content to send.
//
// Returns:
//   - The ID of the sent message.
//   - An error if the operation fails.
func (t *Telegram) Send(msg string) (uint64, error) {
	payload := tgPayload{
		ChatID:    t.chatID,
		Text:      msg,
		ParseMode: "Markdown",
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return 0, fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	var result struct {
		Ok     bool `json:"ok"`
		Result struct {
			MessageID uint64 `json:"message_id"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if !result.Ok {
		return 0, fmt.Errorf("failed to send message")
	}

	if result.Result.MessageID == 0 {
		return 0, fmt.Errorf("failed to get message ID")
	}

	return result.Result.MessageID, nil
}

// Edit modifies an existing message in a Telegram chat.
//
// Parameters:
//   - id: The ID of the message to edit.
//   - msg: The new message content.
//
// Returns:
//   - An error if the operation fails.
func (t *Telegram) Edit(id uint64, msg string) error {
	payload := tgPayload{
		ChatID:    t.chatID,
		MessageID: id,
		Text:      msg,
		ParseMode: "Markdown",
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/editMessageText", t.botToken),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Ok bool `json:"ok"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("failed to edit message")
	}

	return nil
}

// Delete removes an existing message from a Telegram chat.
//
// Parameters:
//   - id: The ID of the message to delete.
//
// Returns:
//   - An error if the operation fails.
func (t *Telegram) Delete(id uint64) error {
	payload := tgPayload{
		ChatID:    t.chatID,
		MessageID: id,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://api.telegram.org/bot%s/deleteMessage", t.botToken),
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Ok bool `json:"ok"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Ok {
		return fmt.Errorf("failed to delete message")
	}

	return nil
}
