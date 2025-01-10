package notify

import (
	"bytes"
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

// Discord represents a client for interacting with Discord webhooks.
type Discord struct {
	webhookID    string
	webhookToken string
}

// NewDiscord creates a new instance of a Discord client.
//
// Parameters:
//   - webhookID: The Discord webhook ID.
//   - webhookToken: The Discord webhook token.
//
// Returns:
//   - A pointer to a Discord instance.
func NewDiscord(webhookID, webhookToken string) *Discord {
	return &Discord{
		webhookID:    webhookID,
		webhookToken: webhookToken,
	}
}

// discordMessage represents the structure of a message sent to Discord via webhook.
type discordMessage struct {
	Content  string `json:"content"`
	Username string `json:"username,omitempty"`
	Avatar   string `json:"avatar_url,omitempty"`
}

// Send sends a message to a Discord channel using the webhook.
//
// Parameters:
//   - msg: The message content to send.
//
// Returns:
//   - The ID of the sent message.
//   - An error if the operation fails.
func (d *Discord) Send(msg string) (uint64, error) {
	discordMsg := discordMessage{
		Content: msg,
	}

	data, err := json.Marshal(discordMsg)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(
		fmt.Sprintf("https://discord.com/api/webhooks/%s/%s?wait=true", d.webhookID, d.webhookToken),
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
		ID string `json:"id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	if result.ID == "" {
		return 0, fmt.Errorf("failed to get message ID")
	}

	// Convert string ID to uint64
	var messageID uint64
	_, err = fmt.Sscanf(result.ID, "%d", &messageID)
	if err != nil {
		return 0, fmt.Errorf("invalid message ID format: %v", err)
	}

	return messageID, nil
}

// Edit modifies an existing message in a Discord channel.
//
// Parameters:
//   - id: The ID of the message to edit.
//   - msg: The new message content.
//
// Returns:
//   - An error if the operation fails.
func (d *Discord) Edit(id uint64, msg string) error {
	discordMsg := discordMessage{
		Content: msg,
	}

	data, err := json.Marshal(discordMsg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("https://discord.com/api/webhooks/%s/%s/messages/%d", d.webhookID, d.webhookToken, id),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to edit message, status code: %d", resp.StatusCode)
	}

	return nil
}

// Delete removes an existing message from a Discord channel.
//
// Parameters:
//   - id: The ID of the message to delete.
//
// Returns:
//   - An error if the operation fails.
func (d *Discord) Delete(id uint64) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://discord.com/api/webhooks/%s/%s/messages/%d", d.webhookID, d.webhookToken, id),
		nil,
	)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("failed to delete message, status code: %d", resp.StatusCode)
	}

	return nil
}
