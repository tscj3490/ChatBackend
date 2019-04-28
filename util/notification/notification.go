package notification

import (
	"fmt"

	"github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

// SendNotification sends
func SendNotification(token string, title string, body string) {
	// To check the token is valid
	pushToken, err := expo.NewExponentPushToken(token)
	if err != nil {
		panic(err)
	}

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       pushToken,
			Body:     body,
			Data:     map[string]string{"withSome": "data"},
			Sound:    "default",
			Title:    title,
			Priority: expo.DefaultPriority,
		},
	)
	// Check errors
	if err != nil {
		panic(err)
		return
	}
	// Validate responses
	if response.ValidateResponse() != nil {
		fmt.Println(response.PushMessage.To, "failed")
	}
}
