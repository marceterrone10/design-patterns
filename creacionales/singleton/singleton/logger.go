package singleton

import "sync"

type Logger struct {
	messages []string
}

var (
	instance *Logger
	once     sync.Once
)

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{
			messages: make([]string, 0),
		}
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.messages = append(l.messages, message)
}

func (l *Logger) GetMessages() []string {
	return l.messages
}
