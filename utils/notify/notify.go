// Package notify provides functionality to send, edit, and delete notifications
// in various messaging platforms, including Discord and Telegram.
package notify

import (
	"fmt"
	"net/http"
)

func closeBody(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		fmt.Printf("Error close response body: %v", err)
	}
}
