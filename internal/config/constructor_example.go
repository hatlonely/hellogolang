package config_test

import (
	"testing"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	_ "github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
)

func Test(t *testing.T) {
	_ = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Username:     "test-user",
		Password:     "123456",
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  200 * time.Millisecond,
		MinIdleConns: 20,
	})

	_, _ = elastic.NewClientFromConfig(&config.Config{
		URL:      "localhost:9200",
		Username: "test-user",
		Password: "123456",
	})

	_, _ = elastic.NewClient(
		elastic.SetURL("localhost:9200"),
		elastic.SetBasicAuth("test-user", "123456"),
	)

	_, _ = oss.New(
		"test-endpoint", "test-ak", "test-sk",
		oss.SecurityToken("test-token"),
		oss.Timeout(1, 20),
	)

	tablestore.NewClient(
		"test-endpoint",
		"test-instance",
		"test-ak",
		"test-sk",
	)
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
}
