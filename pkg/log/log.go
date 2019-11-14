package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// ScooterFormatter represents scooter log formatter
type ScooterFormatter struct {
}

// Format sets log format
func (c *ScooterFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	return []byte(fmt.Sprintf("%s %s scooter[%d]: %s %s\n", timestamp, hostname, os.Getpid(), strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func init() {
	log.SetFormatter(&ScooterFormatter{})
}

// SetLevel sets the log level. Valid levels are panic, fatal, error, warn, info and debug.
func SetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		Fatal(fmt.Sprintf(`not a valid level: "%s"`, level))
	}
	log.SetLevel(lvl)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	log.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
