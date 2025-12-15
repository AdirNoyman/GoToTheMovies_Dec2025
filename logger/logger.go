// Package logger provides a custom logging implementation with separate handlers for info and error logs
package logger

import (
	"log"
	"os"
)

// Logger encapsulates two separate loggers and manages the log file
type Logger struct {
	infoLogger  *log.Logger // *Type = pointer type (stores memory address, not the value itself)
	errorLogger *log.Logger // Pointers allow sharing the same logger across function calls
	file        *os.File    // Must be pointer because os.OpenFile() returns *os.File
}

// NewLogger creates a new logger with output to both file and stdout
// Returns *Logger (pointer) instead of Logger (value) to avoid copying and share state
func NewLogger(logFilePath string) (*Logger, error) {
	// Opens file with: append mode | create if missing | write-only | rw-r--r-- permissions
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// &Logger{...} creates struct instance and returns its memory address (pointer)
	// Without &, would return Logger value instead of *Logger pointer
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		file:        file, // Store file reference for cleanup
	}, nil
}

// Info logs informational messages to stdout
// (l *Logger) = pointer receiver, operates on the original Logger, not a copy
func (l *Logger) Info(msg string) {
	l.infoLogger.Printf("🙄👉 %s", msg)
}

// Error logs error messages to file
// Pointer receiver ensures we write to the actual file, not a copy
func (l *Logger) Error(msg string, err error) {
	l.errorLogger.Printf("%s 😩👉: %v", msg, err)
}

// Close closes the log file to release system resources
// Pointer receiver ensures we close the actual file handle
func (l *Logger) Close() {
	l.file.Close()
}