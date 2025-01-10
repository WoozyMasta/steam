# Notify

The `notify` package provides an easy-to-use interface for sending, editing,
and deleting notifications across multiple messaging platforms, including
Discord and Telegram. It abstracts the underlying API interactions, allowing
developers to integrate messaging functionalities seamlessly into their Go
applications.

## Features

* **Discord Integration**:
  Send, edit, and delete messages via Discord webhooks.
* **Telegram Integration**:
  Send, edit, and delete messages using the Telegram Bot API.
* **Unified Interface**:
  Consistent methods for different platforms.
* **Markdown Support**:
  Utilize Markdown for message formatting where supported.

## Installation

To install the `notify` package, use `go get`:

```bash
go get github.com/woozymasta/steam
```

## Usage

### Discord

```go
package main

import (
  "fmt"
  "log"

  "github.com/woozymasta/steam/utils/notify"
)

func main() {
  // Initialize Discord notifier
  discordWebhookID := "your_discord_webhook_id"
  discordWebhookToken := "your_discord_webhook_token"
  discordClient := notify.NewDiscord(discordWebhookID, discordWebhookToken)

  // Send a message
  messageID, err := discordClient.Send("Hello, **Discord**!")
  if err != nil {
    log.Fatalf("Failed to send Discord message: %v", err)
  }

  // Edit the message
   if err = discordClient.Edit(messageID, "Bye, *Discord*!"); err != nil {
    log.Fatalf("Failed to edit Discord message: %v", err)
  }

  // Delete the message
  if err = discordClient.Delete(messageID); err != nil {
    log.Fatalf("Failed to delete Discord message: %v", err)
  }
}
```

### Telegram

```go
package main

import (
  "fmt"
  "log"

  "github.com/woozymasta/steam/utils/notify"
)

func main() {
  // Initialize Telegram notifier
  telegramBotToken := "your_telegram_bot_token"
  telegramChatID := "your_telegram_chat_id"
  telegramClient := notify.NewTelegram(telegramBotToken, telegramChatID)

  // Send a message
  messageID, err := telegramClient.Send("Hello, *Telegram*!")
  if err != nil {
    log.Fatalf("Failed to send Telegram message: %v", err)
  }

  // Edit the message
  if err = telegramClient.Edit(messageID, "Bye, *Telegram*!"); err != nil {
    log.Fatalf("Failed to edit Telegram message: %v", err)
  }

  // Delete the message
  if err = telegramClient.Delete(messageID); err != nil {
    log.Fatalf("Failed to delete Telegram message: %v", err)
  }
}
```

## Testing

Ensure you have the necessary environment variables set before running tests:

* `DISCORD_ID`: Discord webhook ID.
* `DISCORD_TOKEN`: Discord webhook token.
* `TELEGRAM_ID`: Telegram chat ID.
* `TELEGRAM_TOKEN`: Telegram bot token.

Run the tests using:

```bash
go test ./...
```
