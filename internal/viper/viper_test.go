package viper

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	Convey("test viper", t, func() {
		filename := "test.json"

		So(ioutil.WriteFile(filename, []byte(`{
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
}`), 0644), ShouldBeNil)

		config := viper.New()
		config.AddConfigPath(".")
		config.SetConfigName("test")
		config.SetConfigType("json")

		So(config.ReadInConfig(), ShouldBeNil)

		So(config.GetString("host.address"), ShouldEqual, "localhost")
		So(config.GetInt("datastore.metric.port"), ShouldEqual, 3099)
		So(config.Sub("datastore").GetInt("metric.port"), ShouldEqual, 3099)
		So(config.Sub("abc"), ShouldBeNil)

		So(os.Remove(filename), ShouldBeNil)
	})
}
