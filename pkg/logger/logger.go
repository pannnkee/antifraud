package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func init() {
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}

}
