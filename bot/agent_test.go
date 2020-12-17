package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReply(t *testing.T) {
	chatbot := NewAgent()
	reply := chatbot.Reply("Good Morning")
	assert.Equal(t, "Hey! How are you?", reply)
}

func TestReplyFail(t *testing.T) {
	chatbot := NewAgent()
	reply := chatbot.Reply("")
	assert.Equal(t, "Sorry, I didn't understand.", reply)
}
