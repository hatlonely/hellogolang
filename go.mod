module github.com/hatlonely/hellogolang

go 1.16

//replace (
//	go.uber.org/atomic => github.com/uber-go/atomic v1.4.1-0.20190731194737-ef0d20d85b01
//	go.uber.org/multierr => github.com/uber-go/multierr v1.2.0
//	go.uber.org/zap => github.com/uber-go/zap v1.10.1-0.20190926184545-d8445f34b4ae
//	golang.org/x/net => github.com/golang/net v0.0.0-20190606173856-1492cefac77f
//	golang.org/x/sys => github.com/golang/sys v0.0.0-20190602015325-4c4f7f33c9ed
//	golang.org/x/text => github.com/golang/text v0.3.2
//	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
//	golang.org/x/tools => github.com/golang/tools v0.0.0-20190608022120-eacb66d2a7c3
//	google.golang.org/appengine => github.com/golang/appengine v1.6.1
//	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190605220351-eb0b1bdb6ae6
//	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
//)
//

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/PaesslerAG/gval v1.0.1
	github.com/PaesslerAG/jsonpath v0.1.0
	github.com/ScottMansfield/nanolog v0.2.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/agiledragon/gomonkey v2.0.1+incompatible
	github.com/aliyun/aliyun-mns-go-sdk v0.0.0-20191205082232-b251b9d95415
	github.com/aliyun/aliyun-oss-go-sdk v2.0.7+incompatible
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible
	github.com/antonmedv/expr v1.8.8
	github.com/avast/retry-go v2.6.0+incompatible
	github.com/aws/aws-sdk-go v1.43.21
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/brianvoe/gofakeit/v5 v5.6.2
	github.com/buger/jsonparser v0.0.0-20191004114745-ee4c978eae7e
	github.com/bxcodec/faker/v3 v3.3.1
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/cornelk/hashmap v1.0.1
	github.com/cosmos72/gomacro v0.0.0-20220226114457-23a0d19a6b1e
	github.com/d5/tengo/v2 v2.10.1
	github.com/emirpasic/gods v1.12.0
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/giantswarm/retry-go v0.0.0-20151203102909-d78cea247d5e
	github.com/gin-gonic/gin v1.8.0
	github.com/go-kit/kit v0.9.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gocolly/colly/v2 v2.1.0
	github.com/gogap/errors v0.0.0-20200228125012-531a6449b28c // indirect
	github.com/gogap/stack v0.0.0-20150131034635-fef68dddd4f8 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.5.1
	github.com/hashicorp/consul/api v1.3.0
	github.com/hpifu/go-kit v1.8.0
	github.com/jinzhu/gorm v1.9.15
	github.com/json-iterator/go v1.1.12
	github.com/juju/errgo v0.0.0-20140925100237-08cceb5d0b53 // indirect
	github.com/kpango/glg v1.4.6
	github.com/mailru/easyjson v0.7.7
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	github.com/olivere/elastic/v7 v7.0.32
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/prometheus/client_golang v1.2.1
	github.com/satori/go.uuid v1.2.0
	github.com/sergi/go-diff v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.7.1
	github.com/traefik/yaegi v0.11.3
	github.com/uber/jaeger-client-go v2.24.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7
	github.com/valyala/fasthttp v1.12.0 // indirect
	github.com/yosuke-furukawa/json5 v0.1.1
	github.com/yurishkuro/opentracing-tutorial v0.0.0-20200611023548-a55c44f88513
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20220208050332-20e1d8d225ab // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/protobuf v1.28.0
	gopkg.in/matryer/try.v1 v1.0.0-20150601225556-312d2599e12e
	k8s.io/apimachinery v0.24.3
)
