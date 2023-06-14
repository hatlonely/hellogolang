package aliyun

import (
	"fmt"
	"os"
	"testing"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

type Options struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
}

type Client struct {
	client *openapi.Client
	params *openapi.Params
}

func NewClient(options *Options) (*Client, error) {
	client, err := openapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(options.AccessKeyId),
		AccessKeySecret: tea.String(options.AccessKeySecret),
		Endpoint:        tea.String(options.Endpoint),
	})

	if err != nil {
		return nil, errors.Wrap(err, "openapi.NewClient failed")
	}

	return &Client{
		client: client,
		params: &openapi.Params{
			Action:      tea.String("CreateProject"),
			Version:     tea.String("2020-09-30"),
			Protocol:    tea.String("HTTPS"),
			Method:      tea.String("POST"),
			AuthType:    tea.String("AK"),
			Style:       tea.String("RPC"),
			Pathname:    tea.String("/"),
			ReqBodyType: tea.String("json"),
			BodyType:    tea.String("json"),
		},
	}, nil
}

func (c *Client) Do(body string) (map[string]interface{}, error) {
	return c.client.CallApi(c.params, &openapi.OpenApiRequest{
		Body: body,
	}, &util.RuntimeOptions{
		MaxAttempts: tea.Int(3),
	})
}

func TestCreateProject(t *testing.T) {
	Convey("TestCreateProject", t, func() {
		cli, err := NewClient(&Options{
			AccessKeyId:     os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
			AccessKeySecret: os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
			Endpoint:        "imm.cn-beijing.aliyuncs.com",
		})
		So(err, ShouldBeNil)
		So(cli, ShouldNotBeNil)

		body := `{
			"Action": "CreateProject",
			"ProjectName": "test-project",
		}`

		resp, err := cli.Do(body)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)

		fmt.Print(resp)
	})
}
