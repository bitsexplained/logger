package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
	Critical
)

// LogFormat represents the format of a log message
type LogFormat struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

// LogLocation interface defines methods for logging messages to a location
type LogLocation interface {
	Log(log LogFormat)
}

// LogFile enables creation of a log file
type LogFile struct {
	fileName string
	file     *os.File
	mutex    sync.Mutex
}

// NewFileLog creates a new LogFile based on the given file name
func NewFileLog(fileName string) (*LogFile, error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &LogFile{
		fileName: fileName,
		file:     f,
	}, nil
}

// Log here is used to log on the file
func (f *LogFile) Log(log LogFormat) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	logLine := fmt.Sprintf("[%v:%v] => %v\n", log.Timestamp, log.Level, log.Message)
	f.file.WriteString(logLine)
	defer f.file.Close() // TODO: unsure about the side effect of this
}

// Logger is the configuration object for a log
type Logger struct {
	level       LogLevel
	logFormat   LogFormat
	logLocation []LogLocation
	mutex       sync.Mutex
}

// NewLogger creates a default Logger
func NewLogger() *Logger {
	return &Logger{
		level: Debug,
		logFormat: LogFormat{
			Timestamp: time.Now(),
		},
	}
}

// SetLogLevel sets the logging level
func (l *Logger) SetLogLevel(level LogLevel) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.level = level
}

// SetLogFormat sets the log message format
func (l *Logger) SetLogFormat(logFormat LogFormat) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logFormat = logFormat
}

// AddLogLocation adds a new logging location
func (l *Logger) AddLogLocation(newLocation LogLocation) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logLocation = append(l.logLocation, newLocation)
}

// Log logs a message with the specified level and message to the locations specified
func (l *Logger) Log(level LogLevel, module string, line int, message string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	logFormat := l.logFormat
	logFormat.Level = level
	logFormat.Message = message

	// Log to all specified locations
	for _, location := range l.logLocation {
		location.Log(logFormat)
	}
}
