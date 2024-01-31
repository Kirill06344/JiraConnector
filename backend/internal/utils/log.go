package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"io"
	"os"
)

var Logger = newLogger()

func newLogger() *logrus.Logger {
	logger := logrus.New()

	logFile, err := createLogFile("logs/logs.log")
	if err == nil {
		logger.SetOutput(logFile)
	} else {
		logger.Warnln("Can't create file for path logs/logs.log! Logs will be written in console.")
	}

	errLogFile, err := createLogFile("logs/err_logs.log")
	if err == nil {
		errorWriter := io.MultiWriter(errLogFile, os.Stdout)
		logger.AddHook(&writer.Hook{
			Writer: errorWriter,
			LogLevels: []logrus.Level{
				logrus.WarnLevel,
			},
		})
	} else {
		logger.Warnln("Can't create file for path logs/err_logs.log! Logs will be written in console.")
	}
	return logger
}

func createLogFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}
