package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	LogDebug *logrus.Logger
	LogInfo  *logrus.Logger
	LogError *logrus.Logger
)

func init() {
	LogDebug = createLogger("debug")
	LogInfo = createLogger("info")
	LogError = createLogger("error")
}

func createLogger(logLevel string) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logDir := filepath.Join(".", "logs", logLevel)
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		fmt.Printf("Failed to create log directory: %v\n", err)
		os.Exit(1)
	}

	logFile := filepath.Join(logDir, fmt.Sprintf("%s-%s.log", logLevel, time.Now().Format("2006-01")))

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		os.Exit(1)
	}

	log.SetOutput(file)

	log.AddHook(&formatHook{})

	return log
}

type formatHook struct{}

func (f *formatHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (f *formatHook) Fire(entry *logrus.Entry) error {
	data := make([]interface{}, 0, len(entry.Data))
	for _, v := range entry.Data {
		data = append(data, v)
	}
	entry.Message = fmt.Sprintf(entry.Message, data...)
	return nil
}
