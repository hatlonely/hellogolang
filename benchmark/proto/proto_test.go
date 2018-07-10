package proto

import (
	"encoding/json"
	"testing"

	proto2 "github.com/gogo/protobuf/proto"
	proto1 "github.com/golang/protobuf/proto"
	module2 "github.com/hatlonely/hellogolang/internal/proto/gogoproto"
	module1 "github.com/hatlonely/hellogolang/internal/proto/protobuf"
)

var creative = &module1.Creative{
	Url:             "http://www.baidu.com",
	VideoLength:     123,
	VideoSize:       456,
	VideoResolution: "100*300",
	Width:           11,
	Height:          14,
	WatchMile:       235,
	BitRate:         12,
	IValue:          456,
	SValue:          "ffffffff",
	FValue:          123.456,
	Resolution:      "aaaaaaaa",
	Mime:            "bbbbbbbb",
	AdvCreativeId:   "cccccccc",
	CreativeId:      12567,
	FMd5:            "dddddddd",
	Source:          1,
	Orientation:     2,
	Protocal:        3,
}

var creative2 = &module2.Creative{
	Url:             "http://www.baidu.com",
	VideoLength:     123,
	VideoSize:       456,
	VideoResolution: "100*300",
	Width:           11,
	Height:          14,
	WatchMile:       235,
	BitRate:         12,
	IValue:          456,
	SValue:          "ffffffff",
	FValue:          123.456,
	Resolution:      "aaaaaaaa",
	Mime:            "bbbbbbbb",
	AdvCreativeId:   "cccccccc",
	CreativeId:      12567,
	FMd5:            "dddddddd",
	Source:          1,
	Orientation:     2,
	Protocal:        3,
}

func TestProto(t *testing.T) {
	{
		buf, err := proto1.Marshal(creative)
		err = proto1.Unmarshal(buf, creative)
		t.Log(err, creative)
	}
	{
		buf, err := proto2.Marshal(creative)
		err = proto2.Unmarshal(buf, creative)
		t.Log(err, creative)
	}
	{
		buf1, _ := proto1.Marshal(creative)
		buf2, _ := proto2.Marshal(creative2)
		buf3, _ := json.Marshal(creative)
		t.Log(len(buf1), len(buf2), len(buf3))
	}
	t.Error()
}

func BenchmarkProto(b *testing.B) {
	b.Run("json marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = json.Marshal(creative)
		}
	})
	b.Run("protobuf marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = proto1.Marshal(creative)
		}
	})
	b.Run("gogoproto marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = proto2.Marshal(creative2)
		}
	})

	b.Run("json unmarshal", func(b *testing.B) {
		buf, err := json.Marshal(creative)
		if err != nil {
			panic(err)
		}
		c := &module1.Creative{}
		for i := 0; i < b.N; i++ {
			_ = json.Unmarshal(buf, c)
		}
	})

	buf, err := proto1.Marshal(creative)
	if err != nil {
		panic(err)
	}
	b.Run("protobuf unmarshal", func(b *testing.B) {
		c := &module1.Creative{}
		for i := 0; i < b.N; i++ {
			_ = proto1.Unmarshal(buf, c)
		}
	})
	b.Run("gogoproto unmarshal", func(b *testing.B) {
		c := &module2.Creative{}
		for i := 0; i < b.N; i++ {
			_ = proto2.Unmarshal(buf, c)
		}
	})
}
