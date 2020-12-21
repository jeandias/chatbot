package bot

import watson "github.com/jeandias/chatbot/watson"

type Bot interface {
	Greeting() string
	Reply(string, string) string
}

type Agent struct {
	Bot
}

func (a *Agent) Greeting() string {
	return "Hi!"
}

func (a *Agent) Reply(event string, msg string) string {
	switch event {
	case "welcomeNewUser":
		return "Welcome " + msg + ". Nice to meet you.\n How can I assist you?"

	case "welcomeOldUser":
		return "Welcome back " + msg + ". How can I assist you?"

	case "userChangeName":
		return "You changed your nick name? I liked, " + msg
	default:
		return watson.CallAssistant(msg)
	}
}

func NewAgent() *Agent {
	return &Agent{}
}
