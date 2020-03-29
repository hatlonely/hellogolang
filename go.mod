module github.com/hatlonely/hellogolang

go 1.13

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.0.0-20190603211518-c8433c9aaceb
	go.etcd.io/bbolt => github.com/etcd-io/bbolt v1.3.4-0.20191001164932-6e135e5d7e3d
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.1-0.20190731194737-ef0d20d85b01
	go.uber.org/multierr => github.com/uber-go/multierr v1.2.0
	go.uber.org/zap => github.com/uber-go/zap v1.10.1-0.20190926184545-d8445f34b4ae
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190523035834-f03afa92d3ff
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88
	golang.org/x/net => github.com/golang/net v0.0.0-20190606173856-1492cefac77f
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190602015325-4c4f7f33c9ed
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190608022120-eacb66d2a7c3
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190605220351-eb0b1bdb6ae6
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
)

require (
	github.com/ScottMansfield/nanolog v0.2.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/aliyun/aliyun-oss-go-sdk v2.0.7+incompatible
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible
	github.com/aws/aws-sdk-go v1.25.25
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/buger/jsonparser v0.0.0-20191004114745-ee4c978eae7e
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/cornelk/hashmap v1.0.1
	github.com/emirpasic/gods v1.12.0
	github.com/go-kit/kit v0.9.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/hashicorp/consul/api v1.3.0
	github.com/hpifu/go-kit v1.8.0
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.8
	github.com/kpango/glg v1.4.6
	github.com/mailru/easyjson v0.7.0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/prometheus/client_golang v1.2.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/ugorji/go/codec v1.1.7
	github.com/yosuke-furukawa/json5 v0.1.1
	go.uber.org/zap v1.10.0
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
)
