package aliyun

import (
	"fmt"
	"testing"
	"time"

	"github.com/aliyun/aliyun-mns-go-sdk"
	. "github.com/smartystreets/goconvey/convey"
)

var mnsClient ali_mns.MNSClient

func init() {
	_, accessKeyID, accessKeySecret, err := LoadOSSConfig()
	if err != nil {
		panic(err)
	}
	mnsClient = ali_mns.NewAliMNSClient(
		"http://<user_id>.mns.cn-shanghai.aliyuncs.com/",
		accessKeyID, accessKeySecret)
}

func TestMNSQueue(t *testing.T) {
	Convey("TestMNSQueue", t, func() {
		queue := ali_mns.NewMNSQueue("imm-dev-hl-mns-queue-shanghai", mnsClient)

		{
			res, err := queue.SendMessage(ali_mns.MessageSendRequest{
				MessageBody:  "hello from hatlonely" + time.Now().UTC().Format(time.RFC3339),
				DelaySeconds: 0,
				Priority:     8,
			})
			So(err, ShouldBeNil)
			fmt.Println(res)
		}
		{
			msgChan := make(chan ali_mns.MessageReceiveResponse, 2)
			errChan := make(chan error, 2)
			queue.ReceiveMessage(msgChan, errChan)

			select {
			case msg := <-msgChan:
				fmt.Println(msg.MessageBody)
				So(queue.DeleteMessage(msg.ReceiptHandle), ShouldBeNil)
			case err := <-errChan:
				fmt.Println(err)
			}
		}
	})
}

func TestMNSTopicSubscribe(t *testing.T) {
	Convey("TestMNSTopicSubscribe", t, func() {
		topic := ali_mns.NewMNSTopic("imm-dev-hl-mns-topic-shanghai", mnsClient)
		fmt.Println(topic.GenerateQueueEndpoint("imm-dev-hl-mns-queue-shanghai"))
		err := topic.Subscribe("imm-dev-hl-test-subscript", ali_mns.MessageSubsribeRequest{
			Endpoint:            topic.GenerateQueueEndpoint("imm-dev-hl-mns-queue-shanghai"),
			NotifyContentFormat: ali_mns.SIMPLIFIED,
		})
		So(err, ShouldBeNil)
	})
}

func TestMNSTopic(t *testing.T) {
	Convey("TestMNSTopic", t, func() {
		topic := ali_mns.NewMNSTopic("imm-dev-hl-mns-topic-shanghai", mnsClient)
		topic.PublishMessage(ali_mns.MessagePublishRequest{
			MessageBody: "message from topic" + time.Now().UTC().Format(time.RFC3339),
		})
	})
}
