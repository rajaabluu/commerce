package config

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
	return log
}
