
package pkg

import (
	"log"
	"os"
)

// Logger provides structured logging
type Logger struct {
	*log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile),
	}
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.Printf("[INFO] %s", msg)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.Printf("[ERROR] %s", msg)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	l.Printf("[DEBUG] %s", msg)
}
