package buildin

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	Convey("字符串测试", t, func() {
		Convey("字符串比较", func() {
			name := "hatlonely"
			So(strings.Compare(name, "playjokes"), ShouldEqual, -1)
			So(strings.Compare(name, "hatlonely"), ShouldEqual, 0)
			So(name == "hatlonely", ShouldBeTrue)
		})

		Convey("字符串断言", func() {
			words := "stay hungry, stay foolish"
			So(strings.Contains(words, "hungry"), ShouldBeTrue)
			So(strings.ContainsAny(words, " \t\n"), ShouldBeTrue)
			So(strings.HasPrefix(words, "stay"), ShouldBeTrue)
			So(strings.HasSuffix(words, "foolish"), ShouldBeTrue)
		})

		Convey("字符串查找", func() {
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

		Convey("字符串转化", func() {
			i32, _ := strconv.Atoi("123456")
			So(i32, ShouldEqual, 123456)
			// 10 进制，64位存储
			i64, _ := strconv.ParseInt("123456", 10, 64)
			So(i64, ShouldEqual, 123456)

			So(strconv.Itoa(i32), ShouldEqual, "123456")
			// 10 进制
			So(strconv.FormatInt(i64, 10), ShouldEqual, "123456")

			f64, _ := strconv.ParseFloat("123.456", 64)
			So(f64, ShouldAlmostEqual, 123.456)
			So(strconv.FormatFloat(f64, 'E', -1, 64), ShouldEqual, "1.23456E+02")
			So(fmt.Sprintf("%.3f", f64), ShouldEqual, "123.456")
			So(strconv.FormatBool(true), ShouldEqual, "true")
			// golang 转义
			So(strconv.Quote(`"Fran & Freddie's Diner	☺"`), ShouldEqual, `"\"Fran & Freddie's Diner\t☺\""`)

			So(strconv.QuoteToASCII("Hello, 世界"), ShouldEqual, `"Hello, \u4e16\u754c"`)
		})
	})
}
