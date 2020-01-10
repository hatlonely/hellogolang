package buildin

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGobConvertStruct(t *testing.T) {
	type A struct {
		X int
		Y int
		N string
	}

	type B struct {
		X *int
		Y int64
		N *string
	}

	Convey("test gob convert struct", t, func() {
		buffer := &bytes.Buffer{}
		encoder := gob.NewEncoder(buffer)
		_ = encoder.Encode(&A{1, 2, "s"})

		decoder := gob.NewDecoder(buffer)
		b := &B{}
		So(decoder.Decode(b), ShouldBeNil)
		So(*b.X, ShouldEqual, 1)
		So(b.Y, ShouldEqual, 2)
		So(*b.N, ShouldEqual, "s")
	})
}

type GobV struct {
	x int
	y int
}

func (v GobV) MarshalBinary() ([]byte, error) {
	buffer := &bytes.Buffer{}
	_, _ = fmt.Fprintf(buffer, "%v,%v", v.x, v.y)
	return buffer.Bytes(), nil
}

func (v *GobV) UnmarshalBinary(buf []byte) error {
	buffer := bytes.NewBuffer(buf)
	_, err := fmt.Fscanf(buffer, "%v,%v", &v.x, &v.y)
	return err
}

func TestGobMarshal(t *testing.T) {
	Convey("test marshal", t, func() {
		buffer := &bytes.Buffer{}
		encoder := gob.NewEncoder(buffer)
		_ = encoder.Encode(&GobV{x: 123, y: 456})

		decoder := gob.NewDecoder(buffer)
		v := &GobV{}
		So(decoder.Decode(v), ShouldBeNil)
		So(v.x, ShouldEqual, 123)
		So(v.y, ShouldEqual, 456)
	})
}
