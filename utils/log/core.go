package log

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	Info LogLevel = 0
	Error LogLevel = 1
	Fatal LogLevel = 2
)

func (logLevel LogLevel)ToString() string {
	names := []string{ "Info", "Error", "Fatal" }
	if logLevel < 0 || logLevel > 2 {
		return "unknown"
	}
	return names[logLevel]
}

func (logLevel LogLevel)Print(msg string) {
	var fd *os.File
	
	switch logLevel {
	case Info:
		fd = os.Stdin
	case Error, Fatal:
		fd = os.Stderr
	}

	fmt.Fprintf(fd ,"%s [%s] %s\n", time.Now().Format("2006/01/02 15:04:05"), logLevel.ToString(), msg)
	if logLevel == 2 {
		os.Exit(2)
	}
}
