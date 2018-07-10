package buildin

import (
	"encoding/base64"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBase64(t *testing.T) {
	Convey("test base64", t, func() {
		So(base64.StdEncoding.EncodeToString([]byte("abcdefghijklmnopqrstuvwxyz")), ShouldEqual, "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=")
		So(base64.URLEncoding.EncodeToString([]byte("abcdefghijklmnopqrstuvwxyz")), ShouldEqual, "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=")
		So(base64.RawURLEncoding.EncodeToString([]byte("abcdefghijklmnopqrstuvwxyz")), ShouldEqual, "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo")
		So(base64.RawStdEncoding.EncodeToString([]byte("abcdefghijklmnopqrstuvwxyz")), ShouldEqual, "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo")
	})
}
