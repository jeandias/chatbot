package bot

import (
	"os"
	"testing"

	watson "github.com/jeandias/chatbot/watson"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/jeandias/chatbot/.env"))
	if err != nil {
		panic(err)
	}

	watson.StartWatsonAssistant()
	code := m.Run()
	os.Exit(code)
}

func TestReply(t *testing.T) {
	chatbot := NewAgent()
	reply := chatbot.Reply("", "Hello")
	assert.Equal(t, "{\"output\":"+
		"{\"generic\":[{\"response_type\":\"text\",\"text\":\"Hey! How are you?\"}],"+
		"\"intents\":[{\"intent\":\"action_49172_intent_26844\",\"confidence\":1}]}}", reply)
}

func TestReplyWelcomeOldUser(t *testing.T) {
	chatbot := NewAgent()
	reply := chatbot.Reply("welcomeOldUser", "Bob")
	assert.Equal(t, "Welcome back Bob. How can I assist you?", reply)
}

func TestReplyUserChangeName(t *testing.T) {
	chatbot := NewAgent()
	reply := chatbot.Reply("userChangeName", "Chuck")
	assert.Equal(t, "You changed your nick name? I liked, Chuck", reply)
}
