package main

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
		"oi":  "tudo bem?",
		"sim": "fico feliz em saber que sim",
	}
}

func (a *Agent) Greeting() string {
	return "Olá!"
}

func (a *Agent) Reply(msg string) string {
	return a.knowledgeBase[msg]
}

func NewAgent() *Agent {
	agent := &Agent{}
	agent.initializeIntelligence()
	return agent
}
