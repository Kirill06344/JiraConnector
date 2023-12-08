package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"io"
	"os"
)

func ConfigureLogger() (*log.Logger, error) {
	logger := log.New()

	logFile, err := createLogFile("cmd/logs/logs.log")
	if err != nil {
		return nil, err
	}
	logger.SetOutput(logFile)

	errLogFile, err := createLogFile("cmd/logs/err_logs.log")
	if err != nil {
		return nil, err
	}
	errorWriter := io.MultiWriter(errLogFile, os.Stdout)
	logger.AddHook(&writer.Hook{
		Writer: errorWriter,
		LogLevels: []log.Level{
			log.WarnLevel,
		},
	})
	return logger, nil
}

func createLogFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}
