package seelog

import (
	"testing"
	"github.com/cihub/seelog"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
)

func TestSeelog(t *testing.T) {
	Convey("Given 一个配置文件", t, func() {
		filename := "sample.xml"
		context := `<seelog>
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
		err := ioutil.WriteFile(filename, []byte(context), 0644)
		So(err, ShouldBeNil)

		var woklog seelog.LoggerInterface
		Convey("When 从配置文件中读取", func() {
			woklog, err = seelog.LoggerFromConfigAsFile(filename)
			So(err, ShouldBeNil)
			So(woklog, ShouldNotBeNil)

			Convey("Then 写日志", func() {
				woklog.Infof("Hello seelog!")
			})
		})

		Convey("When 从buffer中读取", func() {
			woklog, err = seelog.LoggerFromConfigAsBytes([]byte(context))
			So(err, ShouldBeNil)
			So(woklog, ShouldNotBeNil)

			Convey("Then 写日志", func() {
				woklog.Infof("Hello seelog!")
			})
		})

		Convey("When 替换默认logger", func() {
			woklog, err = seelog.LoggerFromConfigAsBytes([]byte(context))
			seelog.ReplaceLogger(woklog)
			So(seelog.Current, ShouldEqual, woklog)

			Convey("Then 写日志", func() {
				seelog.Infof("Hello seelog!")
			})
		})

		Convey("Finally 删除文件", func() {
			So(os.Remove(filename), ShouldBeNil)
			So(os.RemoveAll("log"), ShouldBeNil)
		})
	})
}
