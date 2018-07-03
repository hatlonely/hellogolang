package proto

import "testing"
import module1 "github.com/hatlonely/hellogolang/internal/proto/gogoproto"
import module2 "github.com/hatlonely/hellogolang/internal/proto/protobuf"
import proto1 "github.com/gogo/protobuf/proto"
import proto2 "github.com/golang/protobuf/proto"

func TestProto(t *testing.T) {
	{
		creative := &module1.Creative{
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
		buf, err := proto1.Marshal(creative)
		t.Log(buf, err, len(buf))

		err = proto1.Unmarshal(buf, creative)
		t.Log(err, creative)
	}
	{
		creative := &module2.Creative{
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
		buf, err := proto2.Marshal(creative)
		t.Log(buf, err, len(buf))

		err = proto2.Unmarshal(buf, creative)
		t.Log(err, creative)
	}
	t.Error()
}

var creative1 = &module1.Creative{
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

func BenchmarkProto(b *testing.B) {
	b.Run("protobuf marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = proto2.Marshal(creative2)
		}
	})
	b.Run("gogoproto marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = proto1.Marshal(creative1)
		}
	})
	b.Run("protobuf unmarshal", func(b *testing.B) {
		buf, err := proto2.Marshal(creative2)
		if err != nil {
			panic(err)
		}
		creative := &module2.Creative{}
		for i := 0; i < b.N; i++ {
			_ = proto2.Unmarshal(buf, creative)
		}
	})
	b.Run("gogoproto unmarshal", func(b *testing.B) {
		buf, err := proto1.Marshal(creative1)
		if err != nil {
			panic(err)
		}
		creative := &module1.Creative{}
		for i := 0; i < b.N; i++ {
			_ = proto1.Unmarshal(buf, creative)
		}
	})
}
