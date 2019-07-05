package main

import (
	"github.com/guobingithub/logrus-hook/hook"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	apiLogPath := "mock_log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(apiLogPath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(1*time.Minute),
	)
	if err != nil {
		panic("fail to rotatelogs.New, err:" + err.Error())
	}

	writeMap := hook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}

	Log = logrus.New()
	//Log.SetReportCaller(true)
	Log.Hooks.Add(hook.NewBgHook(writeMap, &logrus.JSONFormatter{}))
	return Log
}

func init() {
	Log = NewLogger()
}

func main() {
	for true {
		Log.WithFields(logrus.Fields{
			"name": "tony",
			"age":  18,
			"sex":  1,
		}).Error("A group of walrus emerges from the ocean")
		time.Sleep(5 * time.Second)
	}
}
