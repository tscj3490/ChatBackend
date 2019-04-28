package notification

import (
	"fmt"
	"testing"

	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func TestNotification(t *testing.T) {
	// To check the token is valid
	pushToken, err := expo.NewExponentPushToken("ExponentPushToken[L0KbheBkfgxhIudMN37Okv]")
	if err != nil {
		panic(err)
	}

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       pushToken,
			Body:     "This is a test notification",
			Data:     map[string]string{"withSome": "data"},
			Sound:    "default",
			Title:    "Jacky! Success!",
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
