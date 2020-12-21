package watson

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/assistantv2"
)

var (
	Assistant    *assistantv2.AssistantV2
	AssistantID  string
	SessionID    *string
	assistantErr error
)

func StartWatsonAssistant() (*assistantv2.AssistantV2, string, *string) {
	// Instantiate the Watson AssistantV2 service
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("WATSON_API_KEY"),
	}

	options := &assistantv2.AssistantV2Options{
		Authenticator: authenticator,
		Version:       "2020-12-18",
	}

	Assistant, assistantErr = assistantv2.NewAssistantV2(options)
	// Check successful instantiation
	if assistantErr != nil {
		log.Fatal(assistantErr)
	}

	Assistant.SetServiceURL(os.Getenv("WATSON_API_URL"))

	/* CREATE SESSION */
	// Call the assistant CreateSession method
	AssistantID = os.Getenv("WATSON_ASSISTANT_ID")
	createSessionResult, _, sessionErr := Assistant.CreateSession(
		&assistantv2.CreateSessionOptions{
			AssistantID: core.StringPtr(AssistantID),
		})

	if sessionErr != nil {
		log.Fatal(sessionErr)
	}

	SessionID = createSessionResult.SessionID
	return Assistant, AssistantID, SessionID
}

func CallAssistant(question string) string {
	// 	/* MESSAGE */
	// Call the assistant Message method
	result, _, responseErr := Assistant.Message(
		&assistantv2.MessageOptions{
			AssistantID: core.StringPtr(AssistantID),
			SessionID:   SessionID,
			Input: &assistantv2.MessageInput{
				Text: core.StringPtr(question),
			},
		})

	// Check successful call
	if responseErr != nil {
		DeleteWatsonSession()
		log.Fatal(responseErr)
	}

	b, _ := json.Marshal(result)
	return string(b)
}

func DeleteWatsonSession() {
	// 	/* DELETE SESSION */
	// Call the assistant DeleteSession method
	_, err := Assistant.DeleteSession(&assistantv2.DeleteSessionOptions{
		AssistantID: core.StringPtr(AssistantID),
		SessionID:   SessionID,
	})

	// Check successful call
	if err != nil {
		log.Fatal(err)
	}
}
