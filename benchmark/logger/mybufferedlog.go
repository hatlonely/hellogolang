package logger

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func NewMyBufferedLog(filename string) *MyBufferedLog {
	of, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(of)

	return &MyBufferedLog{
		writer: bw,
	}
}

type MyBufferedLog struct {
	writer *bufio.Writer
	mutex  sync.Mutex
}

func (l *MyBufferedLog) Info(v ...interface{}) {
	l.mutex.Lock()
	l.writer.WriteString(fmt.Sprint(v...))
	l.writer.WriteByte('\n')
	l.mutex.Unlock()
}
