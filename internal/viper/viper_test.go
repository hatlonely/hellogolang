package viper

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	Convey("Given 给一个配置文件", t, func() {
		filename := "test.json"
		context := `{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}`

		err := ioutil.WriteFile(filename, []byte(context), 0644)
		So(err, ShouldBeNil)

		Convey("When 读取配置文件", func() {
			config := viper.New()

			config.AddConfigPath(".")
			config.SetConfigName("test")
			config.SetConfigType("json")

			err := config.ReadInConfig()
			So(err, ShouldBeNil)

			So(config.GetString("host.address"), ShouldEqual, "localhost")
			So(config.GetInt("datastore.metric.port"), ShouldEqual, 3099)
			So(config.Sub("abc"), ShouldBeNil)
		})

		Convey("Finally 删除文件", func() {
			os.Remove(filename)
		})
	})
}
