package aws

import (
	"bufio"
	"context"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	. "github.com/smartystreets/goconvey/convey"
)

func TestS3(t *testing.T) {
	Convey("连接到s3", t, func() {
		sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String(endpoints.ApSoutheast1RegionID),
		}))
		service := s3.New(sess)

		Convey("上传文件", func() {
			fp, err := os.Open("s3_test.go")
			So(err, ShouldBeNil)
			defer fp.Close()

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
			defer cancel()

			_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
				Bucket: aws.String("hatlonely"),
				Key:    aws.String("test/s3_test.go"),
				Body:   fp,
			})
			So(err, ShouldBeNil)
		})

		Convey("下载文件", func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
			defer cancel()

			out, err := service.GetObjectWithContext(ctx, &s3.GetObjectInput{
				Bucket: aws.String("hatlonely"),
				Key:    aws.String("test/s3_test.go"),
			})
			So(err, ShouldBeNil)
			defer out.Body.Close()
			scanner := bufio.NewScanner(out.Body)
			for scanner.Scan() {
				Println(scanner.Text())
			}
		})

		Convey("遍历目录 ListObjectsPages", func() {
			var objkeys []string

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
			defer cancel()

			err := service.ListObjectsPagesWithContext(ctx, &s3.ListObjectsInput{
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

		Convey("遍历目录 ListObjects", func() {
			var objkeys []string

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
			defer cancel()

			out, err := service.ListObjectsWithContext(ctx, &s3.ListObjectsInput{
				Bucket: aws.String("hatlonely"),
				Prefix: aws.String("test/"),
			})
			So(err, ShouldBeNil)
			for _, content := range out.Contents {
				objkeys = append(objkeys, aws.StringValue(content.Key))
			}
			Println(objkeys)
		})
	})
}
