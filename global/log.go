package global

import "github.com/juju/loggo"

var log = loggo.GetLogger("Toolkit-service")

func GetLogger() loggo.Logger {
	return log
}

func SetLevel(level loggo.Level) {
	log.SetLogLevel(level)
}
