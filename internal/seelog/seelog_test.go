package seelog

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSeelog(t *testing.T) {
	Convey("test seelog", t, func() {
		config := `<seelog>
    <outputs>
        <filter levels="debug, trace, info, warn, error, critical">
            <buffered formatid="runtime" size="10000" flushperiod="100">
                <rollingfile type="date" filename="./log/sample.woklog" datepattern="2006010215" maxrolls="240"/>
            </buffered>
        </filter>
    </outputs>
    <formats>
        <format id="runtime" format="%Date %Time [%Level] [%RelFile:%Line] [%Func] %Msg%n"/>
    </formats>
</seelog>`

		var err error
		var woklog seelog.LoggerInterface
		Convey("load from file", func() {
			So(ioutil.WriteFile("sample.xml", []byte(config), 0644), ShouldBeNil)
			woklog, err = seelog.LoggerFromConfigAsFile("sample.xml")
			So(err, ShouldBeNil)
			So(woklog, ShouldNotBeNil)
			woklog.Infof("Hello seelog!")
			woklog.Close()
			So(os.Remove("sample.xml"), ShouldBeNil)
			So(os.RemoveAll("log"), ShouldBeNil)
		})

		Convey("load from byte", func() {
			woklog, err = seelog.LoggerFromConfigAsBytes([]byte(config))
			So(err, ShouldBeNil)
			So(woklog, ShouldNotBeNil)
			woklog.Infof("Hello seelog!")
		})

		Convey("replace logger", func() {
			woklog, err = seelog.LoggerFromConfigAsBytes([]byte(config))
			seelog.ReplaceLogger(woklog)
			So(seelog.Current, ShouldEqual, woklog)
			seelog.Infof("Hello seelog!")
		})
	})
}
