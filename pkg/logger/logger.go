package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var Log = logrus.New()

func InitLogger() {
	// Log output to file & console
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)
	}

	// Set log format
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// Set log level (INFO, WARN, ERROR, DEBUG)
	Log.SetLevel(logrus.DebugLevel)
}
