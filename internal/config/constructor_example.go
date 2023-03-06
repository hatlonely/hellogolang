package config_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	_ "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
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

}
