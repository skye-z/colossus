package main

import (
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type FileLogger struct {
	filename string
}

func NewFileLogger(filename string) logger.Logger {
	return &FileLogger{
		filename: filename,
	}
}

func (l *FileLogger) Print(message string) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Print(message)
}

func (l *FileLogger) Println(message string) {
	l.Print(message + "\n")
}

func (l *FileLogger) Trace(message string) {
	l.Println("TRACE\t" + message)
}

func (l *FileLogger) Debug(message string) {
	l.Println("DEBUG\t" + message)
}

func (l *FileLogger) Info(message string) {
	l.Println("INFO \t" + message)
}

func (l *FileLogger) Warning(message string) {
	l.Println("WARN \t" + message)
}

func (l *FileLogger) Error(message string) {
	l.Println("ERROR\t" + message)
}

func (l *FileLogger) Fatal(message string) {
	l.Println("FATAL\t" + message)
	os.Exit(1)
}
