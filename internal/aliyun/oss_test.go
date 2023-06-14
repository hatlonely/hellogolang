package aliyun

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	. "github.com/smartystreets/goconvey/convey"
)

var ossClient *oss.Client

func init() {
	var err error
	accessKeyID := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
	ossClient, err = oss.New("oss-cn-beijing.aliyuncs.com", accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
}

func TestBucket(t *testing.T) {
	Convey("test list bucket", t, func() {
		res, err := ossClient.ListBuckets()
		So(err, ShouldBeNil)

		for _, bucket := range res.Buckets {
			_, _ = Println(bucket.Name)
		}
	})
}

func TestObject(t *testing.T) {
	Convey("test object", t, func() {
		bucket, err := ossClient.Bucket("hatlonely-test-bucket")
		So(err, ShouldBeNil)

		Convey("put object", func() {
			err := bucket.PutObjectFromFile("oss_test.go", "oss_test.go")
			So(err, ShouldBeNil)
		})

		Convey("get object", func() {
			fp, err := bucket.GetObject("oss_test.go")
			So(err, ShouldBeNil)
			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				_, _ = Println(scanner.Text())
			}
		})

		Convey("list object", func() {
			res, err := bucket.ListObjects()
			So(err, ShouldBeNil)

			for _, obj := range res.Objects {
				fmt.Println(obj.Key)
			}
		})

		Convey("del object", func() {
			err := bucket.DeleteObject("oss_test.go")
			So(err, ShouldBeNil)
		})
	})
}
