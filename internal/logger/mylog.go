package logger

import (
	"fmt"
	"os"
)

func NewMyLog(filename string) *MyLog {
	of, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return &MyLog{
		fp: of,
	}
}

type MyLog struct {
	fp *os.File
}

func (l *MyLog) Info(v ...interface{}) {
	l.fp.WriteString(fmt.Sprint(v...))
	l.fp.WriteString("\n")
}
