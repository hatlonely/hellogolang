package js5

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

func TestJson5(t *testing.T) {
	Convey("test json5", t, func() {
		str := `{
			"address": "127.0.0.1:8500",
			"service": "add",
			"tag": [
				"tag1", "tag2"
			],
			// "deregisterCriticalServiceAfter": "1m",
			// "interval": "10s"
			"port": 3000
		}`

		var v interface{}
		err := json5.Unmarshal([]byte(str), &v)
		So(err, ShouldBeNil)

		type MyStruct struct {
			Address                        string
			Service                        string
			Tag                            []string
			Port                           int
			DeregisterCriticalServiceAfter time.Duration
			Interval                       time.Duration
		}
		var ms MyStruct
		err = json5.Unmarshal([]byte(str), &ms)
		So(err, ShouldBeNil)
		So(ms.Address, ShouldEqual, "127.0.0.1:8500")
		So(ms.Service, ShouldEqual, "add")
		So(ms.Tag, ShouldResemble, []string{"tag1", "tag2"})
		So(ms.Port, ShouldEqual, 3000)
	})
}
