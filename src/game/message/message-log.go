package message

import "github.com/gookit/event"

type Message struct {
	Time    string
	Content string
}

type MessageLog struct {
	messages []Message
}

func (l *MessageLog) AddMessage(message Message) {
	l.messages = append(l.messages, message)
}

func (l *MessageLog) GetMessages() []Message {
	return l.messages
}

func (l *MessageLog) AddMessageLogSubscriber(e event.Event) error {
	l.AddMessage(e.Get("message").(Message))

	return nil
}
