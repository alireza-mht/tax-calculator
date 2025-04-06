package log

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// InitLogger setups the logger
func InitLogger(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Fatalf("Invalid log level: %s", level)
	}
	log = logrus.New()
	log.SetLevel(logLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func GetLogger() *logrus.Logger {
	return log
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
