package logger

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ScottMansfield/nanolog"
)

// 使用上很不友好，不如直接用 bufio
func NewNanoLog(filename string) *NanoLog {
	log := nanolog.New()
	h := log.AddLogger("%s\n")
	of, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(of)
	log.SetWriter(bw)
	return &NanoLog{
		log: log,
		h:   h,
	}
}

type NanoLog struct {
	log nanolog.LogWriter
	h   nanolog.Handle
}

func (l *NanoLog) Info(v ...interface{}) {
	l.log.Log(l.h, fmt.Sprint(v...)+"\n")
}
