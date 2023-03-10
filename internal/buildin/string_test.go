package buildin

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("test string", t, func() {
		Convey("string compare", func() {
			name := "hatlonely"
			So(strings.Compare(name, "playjokes"), ShouldEqual, -1)
			So(strings.Compare(name, "hatlonely"), ShouldEqual, 0)
			So(name == "hatlonely", ShouldBeTrue)
		})

		Convey("string predicate", func() {
			words := "stay hungry, stay foolish"
			So(strings.Contains(words, "hungry"), ShouldBeTrue)
			So(strings.ContainsAny(words, " \t\n"), ShouldBeTrue)
			So(strings.HasPrefix(words, "stay"), ShouldBeTrue)
			So(strings.HasSuffix(words, "foolish"), ShouldBeTrue)
		})

		Convey("string find", func() {
			// unicode 编码，一个中文占三个字符，数字和字母占一个字符，可以 rune 修饰中文字符
			words := "遇见你，就好像见到了120斤的运气"
			So(len(words), ShouldEqual, 45)
			So(strings.Index(words, "运气"), ShouldEqual, 39)
			So(strings.Index(words, "hello"), ShouldEqual, -1)
			So(strings.IndexAny(words, "123456"), ShouldEqual, 30)
			So(strings.IndexRune(words, rune('运')), ShouldEqual, 39)
			So(strings.LastIndex(words, "运气"), ShouldEqual, 39)
		})

		Convey("字符串操作", func() {
			So("hello"+"world", ShouldEqual, "helloworld")
			So(strings.Split("golang swift", " "), ShouldResemble, []string{"golang", "swift"})
			So(strings.Join([]string{"golang", "swift"}, " "), ShouldEqual, "golang swift")
			So(strings.ToUpper("hello"), ShouldEqual, "HELLO")
			So(strings.ToLower("World"), ShouldEqual, "world")
			So(strings.Trim(" hello world ", " \t\n"), ShouldEqual, "hello world")
			So(strings.TrimLeft(" hello world ", " \t\n"), ShouldEqual, "hello world ")
			So(strings.TrimRight(" hello world ", " \t\n"), ShouldEqual, " hello world")
			So(strings.TrimSpace(" hello world "), ShouldEqual, "hello world")
			So(strings.Repeat("na", 3), ShouldEqual, "nanana")
			So(strings.Replace("oink oink oink", "oink", "moo", -1), ShouldEqual, "moo moo moo")
			So(strings.Replace("oink oink oink", "oink", "moo", 2), ShouldEqual, "moo moo oink")

			So("0123456789"[3:6], ShouldEqual, "345")
			So("0123456789"[:6], ShouldEqual, "012345")
			So("0123456789"[7:], ShouldEqual, "789")
		})

		Convey("string convert", func() {
			b, _ := strconv.ParseBool("true")
			So(b, ShouldEqual, true)

			i, _ := strconv.Atoi("123456")
			So(i, ShouldEqual, 123456)

			i8, _ := strconv.ParseInt("123", 10, 8)
			So(i8, ShouldEqual, 123)

			i16, _ := strconv.ParseInt("12345", 10, 16)
			So(i16, ShouldEqual, 12345)

			i32, _ := strconv.ParseInt("1234567890", 10, 32)
			So(i32, ShouldEqual, 1234567890)

			i64, _ := strconv.ParseInt("123456789123456789", 10, 64)
			So(i64, ShouldEqual, 123456789123456789)

			u8, _ := strconv.ParseUint("123", 10, 8)
			So(u8, ShouldEqual, 123)

			u16, _ := strconv.ParseUint("12345", 10, 16)
			So(u16, ShouldEqual, 12345)

			u32, _ := strconv.ParseUint("1234567890", 10, 32)
			So(u32, ShouldEqual, 1234567890)

			u64, _ := strconv.ParseUint("123456789123456789", 10, 64)
			So(u64, ShouldEqual, 123456789123456789)

			f32, _ := strconv.ParseFloat("123.456", 10)
			So(f32, ShouldAlmostEqual, 123.456)

			f64, _ := strconv.ParseFloat("123.456", 10)
			So(f64, ShouldAlmostEqual, 123.456)

			So(strconv.Itoa(123), ShouldEqual, "123")
			So(strconv.FormatInt(123456, 10), ShouldEqual, "123456")
			So(strconv.FormatUint(123456, 10), ShouldEqual, "123456")
			So(strconv.FormatBool(true), ShouldEqual, "true")
			So(strconv.FormatFloat(123.456, 'E', -1, 64), ShouldEqual, "1.23456E+02")
			So(strconv.FormatFloat(123.456, 'f', -1, 64), ShouldEqual, "123.456")
			So(fmt.Sprintf("%.3f", f64), ShouldEqual, "123.456")
		})

		Convey("quote", func() {
			// golang 转义
			So(strconv.Quote(`"Fran & Freddie's Diner	☺"`), ShouldEqual, `"\"Fran & Freddie's Diner\t☺\""`)
			So(strconv.QuoteToASCII("Hello, 世界"), ShouldEqual, `"Hello, \u4e16\u754c"`)
		})

		Convey("filepath", func() {
			path := "/home/work/hello.docx"
			So(filepath.Dir(path), ShouldEqual, "/home/work")
			So(filepath.Base(path), ShouldEqual, "hello.docx")
			So(filepath.Ext(path), ShouldEqual, ".docx")
		})
	})
}
