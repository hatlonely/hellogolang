package logger

import (
	"bufio"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogrusLog(filename string) *LogrusLog {
	of, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(of)
	log := logrus.New()
	log.Out = bw
	return &LogrusLog{
		log: log,
	}
}

type LogrusLog struct {
	log *logrus.Logger
}

func (l *LogrusLog) Info(v ...interface{}) {
	l.log.Info(v...)
}
