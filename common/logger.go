package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Loginit 初始化日志
func LogToFile() {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	logFile := "./log/" + date + ".log"
	if file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}
}

// Log 记录日志
func LogError(ns string, log interface{}) {
	logrus.Error(fmt.Sprintf("[%s],%+v", ns, log))
}
func LogInfo(ns string, log interface{}) {
	logrus.Info(fmt.Sprintf("[%s],%+v", ns, log))
}

func LogDebug(ns string, log interface{}) {
	logrus.Debugf(fmt.Sprintf("[%s],%s", ns, log))
}
