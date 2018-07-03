package logger

import (
	"bufio"
	"log"
	"os"
)

func NewStdLog(filename string) *StdLog {
	of, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(of)
	log.SetOutput(bw)

	return &StdLog{}
}

type StdLog struct {
}

func (l *StdLog) Info(v ...interface{}) {
	log.Print(v...)
}
