package logger

import (
	"fmt"

	"github.com/kpango/glg"
)

// 这个库有点问题，Info 不好使，FileWriter 也不好使
func NewGlgLog(filename string) *GlgLog {
	fw := glg.FileWriter(filename, 0666)
	log := glg.New()
	log.SetWriter(fw)
	return &GlgLog{
		log: log,
	}
}

type GlgLog struct {
	log *glg.Glg
}

func (l *GlgLog) Info(v ...interface{}) {
	l.log.Info(fmt.Sprint(v...))
}
