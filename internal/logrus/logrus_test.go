package logrus

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func TestLogrus(t *testing.T) {
	Convey("一般用法", t, func() {
		Convey("text formatter", func() {
			log := logrus.New()
			log.Out = os.Stdout
			textFormatter := &logrus.TextFormatter{}
			log.Formatter = textFormatter
			log.WithFields(logrus.Fields{"animal": "walrus", "size": 10}).Infof("A group of walrus emerges from the ocean")
		})

		Convey("json formatter", func() {
			log := logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{}
			log.WithFields(logrus.Fields{"animal": "walrus", "size": 10}).Infof("A group of walrus emerges from the ocean")
		})

		Convey("添加文件和行号", func() {
			log := logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{}
			_, file, line, _ := runtime.Caller(0)
			log.WithFields(logrus.Fields{"file": file, "line": line}).Infof("hello world")
		})
	})
}
