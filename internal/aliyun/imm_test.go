package aliyun

import (
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
}

type Client struct {
	client *openapi.Client
	params *openapi.Params
}

func NewClient(options *Options) (*Client, error) {
	client, err := openapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(options.AccessKeyId),
		AccessKeySecret: tea.String(options.AccessKeySecret),
		Endpoint:        tea.String("imm-cn-beijing.aliyuncs.com"),
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
	}, &util.RuntimeOptions{MaxAttempts: tea.Int(3)})
}

func TestCreateProject(t *testing.T) {
	Convey("TestCreateProject", t, func() {
	cli, err := NewClient(&Options{
	AccessK
	})
}

/**
 * API 相关
 * @param path params
 * @return OpenApi.Params
 */
func CreateApiInfo() (_result *openapi.Params) {
	params := &openapi.Params{
		// 接口名称
		Action: tea.String("CreateProject"),
		// 接口版本
		Version: tea.String("2020-09-30"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("RPC"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func _main(args []*string) (_err error) {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例使用环境变量获取 AccessKey 的方式进行调用，仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, _err := CreateClient(tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")), tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")))
	if _err != nil {
		return _err
	}

	params := CreateApiInfo()
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	_, _err = client.CallApi(params, request, runtime)
	if _err != nil {
		return _err
	}
	return _err
}

//func main() {
//	err := _main(tea.StringSlice(os.Args[1:]))
//	if err != nil {
//		panic(err)
//	}
//}
