package config_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func Test(t *testing.T) {
	// NewClient(opt *Options) *Client
	_ = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Username:     "test-user",
		Password:     "123456",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  200 * time.Millisecond,
		MinIdleConns: 20,
	})

	// NewClientFromConfig(cfg *config.Config) (*Client, error)
	_, _ = elastic.NewClientFromConfig(&config.Config{
		URL:      "localhost:9200",
		Username: "test-user",
		Password: "123456",
	})

	// NewClient(options ...ClientOptionFunc) (*Client, error)
	_, _ = elastic.NewClient(
		elastic.SetURL("localhost:9200"),
		elastic.SetBasicAuth("test-user", "123456"),
	)

	// New(endpoint, accessKeyID, accessKeySecret string, options ...ClientOption) (*Client, error)
	_, _ = oss.New(
		"test-endpoint", "test-ak", "test-sk",
		oss.SecurityToken("test-token"),
		oss.Timeout(1, 20),
	)

	// NewClient(endPoint, instanceName, accessKeyId, accessKeySecret string, options ...ClientOption) *TableStoreClient
	tablestore.NewClient(
		"test-endpoint",
		"test-instance",
		"test-ak",
		"test-sk",
	)

	// NewClientWithConfig(endPoint, instanceName, accessKeyId, accessKeySecret string, securityToken string, config *TableStoreConfig) *TableStoreClient
	tablestore.NewClientWithConfig(
		"test-endpoint",
		"test-instance",
		"test-ak",
		"test-sk",
		"test-token",
		&tablestore.TableStoreConfig{
			RetryTimes:   3,
			MaxRetryTime: 10 * time.Second,
			HTTPTimeout: tablestore.HTTPTimeout{
				ConnectionTimeout: 200 * time.Minute,
			},
			MaxIdleConnections: 200,
		},
	)

	// NewClientWithOptions(regionId string, config *sdk.Config, credential auth.Credential) (client *Client, err error)
	_, _ = sts.NewClientWithOptions(
		"cn-beijing",
		&sdk.Config{
			AutoRetry:    true,
			MaxRetryTime: 2,
			UserAgent:    "",
			Debug:        false,
			Scheme:       "http",
			Timeout:      5 * time.Second,
		},
		credentials.AccessKeyCredential{
			AccessKeyId:     "test-ak",
			AccessKeySecret: "test-sk",
		},
	)

	// NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret string) (client *Client, err error)
	_, _ = sts.NewClientWithAccessKey(
		"cn-beijing",
		"test-ak",
		"test-sk",
	)

	// NewClientWithEcsRamRole(regionId string, roleName string) (client *Client, err error)
	_, _ = sts.NewClientWithEcsRamRole(
		"cn-beijing",
		"test-role",
	)

	// Open(driverName, dataSourceName string) (*DB, error)
	mySqlCli, _ := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"test-user", "123456", "localhost", 3306, "test-db",
	))
	mySqlCli.SetMaxIdleConns(10)
	mySqlCli.SetMaxOpenConns(20)
	mySqlCli.SetConnMaxLifetime(10 * time.Minute)

	// NewServer(opt ...ServerOption) *Server
	_ = grpc.NewServer(
		grpc.ConnectionTimeout(time.Second),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 500 * time.Millisecond,
			MaxConnectionAge:  5 * time.Second,
			Timeout:           4 * time.Second,
		}),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			return handler(ctx, req)
		}),
	)

	// New() *Engine
	ginEngine := gin.New()
	ginEngine.UseRawPath = true
	ginEngine.RedirectTrailingSlash = true
	ginEngine.HandleMethodNotAllowed = true

	// NewServeMux(opts ...ServeMuxOption) *ServeMux
	runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
			if strings.HasPrefix(s, "x-") {
				return s, true
			}
			return "", false
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
		}),
	)
}
