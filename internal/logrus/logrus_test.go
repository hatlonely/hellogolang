package logrus

import (
	"encoding/json"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

type Person struct {
	Name     string              `json:"name"`
	Birthday time.Time           `json:"birthday"`
	Emails   []string            `json:"emails"`
	Skills   map[string]struct{} `json:"skill"`
	Other    string              `json:"-"`
}

func TestJson(t *testing.T) {
	Convey("Json 的用法", t, func() {
		birthday, err := time.Parse("2006-01-02", "2018-03-24")
		So(err, ShouldBeNil)

		person := &Person{
			Name:     "hatlonely",
			Birthday: birthday,
			Emails:   []string{"hatlonely@foxmail.com", "hatlonely@gmail.com"},
			Skills:   map[string]struct{}{"golang": {}, "java": {}, "c++": {}},
			Other:    "hello world",
		}

		buf, err := json.Marshal(person)
		So(err, ShouldBeNil)
		So(string(buf), ShouldEqual, `{"name":"hatlonely","birthday":"2018-03-24T00:00:00Z","emails":["hatlonely@foxmail.com","hatlonely@gmail.com"],"skill":{"c++":{},"golang":{},"java":{}}}`)

		log := logrus.New()
		log.Out = os.Stdout
		log.Formatter = &logrus.JSONFormatter{}
		log.WithFields(logrus.Fields{"person": person}).Error()
	})
}

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
