package logger

import (
	"bufio"
	"fmt"
	"os"

	logging "github.com/op/go-logging"
)

// 这个库使用上不是很友好，设计得有点奇怪
func NewOpLoggingLog(filename string) *OpLoggingLog {
	of, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(of)

	log := logging.MustGetLogger("example")
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	be := logging.NewLogBackend(bw, "", 0)
	bef := logging.NewBackendFormatter(be, format)
	bel := logging.AddModuleLevel(be)
	logging.SetBackend(bel, bef)

	return &OpLoggingLog{
		log: log,
	}
}

type OpLoggingLog struct {
	log *logging.Logger
}

func (l *OpLoggingLog) Info(v ...interface{}) {
	l.log.Info(fmt.Sprint(v))
}
