package utils

import (
    "github.com/sirupsen/logrus"
    "os"
)

func NewLogger() *logrus.Logger {
    logger := logrus.New()
    logger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    logger.SetOutput(os.Stdout)
    return logger
}
