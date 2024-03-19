package log

import (
	"github.com/sirupsen/logrus"
)

type Logger = logrus.FieldLogger

func Default() Logger {
	return logrus.StandardLogger()
}
