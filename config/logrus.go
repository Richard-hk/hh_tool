package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() {
	infoLogPath := "log/info.log"
	infoFile, _ := os.OpenFile(infoLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetOutput(infoFile)
}
