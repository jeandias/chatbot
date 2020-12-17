package bot

import "strings"

type Bot interface {
	Greeting() string
	Reply(string) string
}

type Agent struct {
	Bot
	knowledgeBase map[string]string
}

func (a *Agent) initializeIntelligence() {
	a.knowledgeBase = map[string]string{
		"hey":            "Hey! How are you?",
		"hello":          "Hey! How are you?",
		"hi":             "Hey! How are you?",
		"good morning":   "Hey! How are you?",
		"good evening":   "Hey! How are you?",
		"hey there":      "Hey! How are you?",
		"yo":             "Hey! How are you?",
		"bye":            "Goodbye!",
		"goodbye":        "Goodbye!",
		"see you around": "Goodbye!",
		"see you later":  "Goodbye!",
	}
}

func (a *Agent) Greeting() string {
	return "Can I help you?"
}

func (a *Agent) Reply(msg string) string {
	key := strings.TrimSpace(strings.ToLower(msg))
	reply := a.knowledgeBase[key]

	if reply == "" {
		reply = "Sorry, I didn't understand."
	}
	return reply
}

func NewAgent() *Agent {
	agent := &Agent{}
	agent.initializeIntelligence()
	return agent
}
