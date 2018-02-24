package aws_sdk_go

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"context"
	"bufio"
)

func TestS3(t *testing.T) {
	Convey("连接到s3", t, func() {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
		}))
		svc := s3.New(sess)

		Convey("上传文件", func() {
			fp, err := os.Open("s3_test.go")
			defer fp.Close()

			So(err, ShouldBeNil)
			_, err = svc.PutObjectWithContext(context.Background(), &s3.PutObjectInput{
				Bucket: aws.String("hatlonely"),
				Key:    aws.String("test/s3_test.go"),
				Body:   fp,
			})
			So(err, ShouldBeNil)
		})

		Convey("下载文件", func() {
			out, err := svc.GetObjectWithContext(context.Background(), &s3.GetObjectInput{
				Bucket: aws.String("hatlonely"),
				Key: aws.String("test/s3_test.go"),
			})
			So(err, ShouldBeNil)
			defer out.Body.Close()
			scanner := bufio.NewScanner(out.Body)
			for scanner.Scan() {
				Println(scanner.Text())
			}
		})

		Convey("遍历目录", func() {
			var objkeys []string
			err := svc.ListObjectsPagesWithContext(context.Background(), &s3.ListObjectsInput{
				Bucket: aws.String("hatlonely"),
				Prefix: aws.String("test/"),
			}, func(output *s3.ListObjectsOutput, b bool) bool {
				for _, content := range output.Contents {
					objkeys = append(objkeys, aws.StringValue(content.Key))
				}
				return true
			})
			So(err, ShouldBeNil)
			Println(objkeys)
		})
	})
}
