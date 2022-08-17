package buildin

import (
	"regexp"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegex(t *testing.T) {
	Convey("test regex", t, func() {
		Convey("test new regex", func() {
			{
				pattern, err := regexp.Compile("^[a-z0-9]+@[a-z0-9.]+[.][a-z]{2,4}$")
				So(err, ShouldBeNil)
				So(pattern, ShouldNotBeNil)
			}
			{
				So(regexp.MustCompile("^[a-z0-9]+@[a-z0-9.]+[.][a-z]{2,4}$"), ShouldNotBeNil)
			}
			{
				pattern, err := regexp.CompilePOSIX("^[a-z0-9]+@[a-z0-9.]+[.][a-z]{2,4}$")
				So(err, ShouldBeNil)
				So(pattern, ShouldNotBeNil)
			}
			{
				So(regexp.MustCompilePOSIX("^[a-z0-9]+@[a-z0-9.]+[.][a-z]{2,4}$"), ShouldNotBeNil)
			}
		})

		Convey("test match", func() {
			{
				So(regexp.MustCompile("hello").Match([]byte("hello world")), ShouldBeTrue)
				So(regexp.MustCompile("hello").MatchString("hello world"), ShouldBeTrue)
				So(regexp.MustCompile("hello").MatchReader(strings.NewReader("hello world")), ShouldBeTrue)
			}
			{
				ok, err := regexp.MatchString("hello", "hello world")
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			}
			{
				ok, err := regexp.Match("hello", []byte("hello world"))
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			}
			{
				ok, err := regexp.MatchReader("hello", strings.NewReader("hello world"))
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			}
		})

		Convey("test wildcard", func() {
			So(regexp.MustCompile("[0-9]*").MatchString(""), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]?").MatchString(""), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]?").MatchString("1"), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]+").MatchString("123"), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]{3}").MatchString("123"), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]{3,4}").MatchString("123"), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]{3,4}").MatchString("1234"), ShouldBeTrue)
			So(regexp.MustCompile("[0-9]{3,}").MatchString("12345"), ShouldBeTrue)
			So(regexp.MustCompile("\\d+").MatchString("123"), ShouldBeTrue)
			So(regexp.MustCompile("\\s+").MatchString(" \t"), ShouldBeTrue)
			So(regexp.MustCompile("\\w+").MatchString("abc"), ShouldBeTrue)
			So(regexp.MustCompile("(f|z)oo").MatchString("foo"), ShouldBeTrue)
			So(regexp.MustCompile("(f|z)oo").MatchString("zoo"), ShouldBeTrue)
			So(regexp.MustCompile(".*").MatchString("any string"), ShouldBeTrue)
		})

		Convey("test capture", func() {
			{
				pattern := regexp.MustCompile("[a-z0-9]+@[a-z0-9.]+[.][a-z]{2,4}")
				So(pattern.FindString(" hatlonely@foxmail.com "), ShouldEqual, "hatlonely@foxmail.com")
			}
			{
				pattern := regexp.MustCompile("([a-z0-9]+)@(([a-z0-9.]+)[.]([a-z]{2,4}))")
				So(pattern.FindStringSubmatch(" hatlonely@foxmail.com "), ShouldResemble, []string{
					"hatlonely@foxmail.com",
					"hatlonely",
					"foxmail.com",
					"foxmail",
					"com",
				})
			}
		})

		Convey("not support predict", func() {
			_, err := regexp.Compile("Windows (?=95|98|NT|2000)")
			So(err, ShouldNotBeNil)
		})

		Convey("not support back reference", func() {
			_, err := regexp.Compile("(\\w+) \\1")
			So(err, ShouldNotBeNil)
		})

		Convey("test replace", func() {
			pattern := regexp.MustCompile("^([a-z0-9]+)@(?:([a-z0-9.]+)[.]([a-z]{2,4}))$")
			So(pattern.ReplaceAllString("hatlonely@foxmail.com", "$0 $1 $2 $3"), ShouldEqual, "hatlonely@foxmail.com hatlonely foxmail com")
		})

		Convey("test split", func() {
			pattern := regexp.MustCompile("\\s+")
			So(pattern.Split("abab  x   acac y adad", -1), ShouldResemble, []string{
				"abab", "x", "acac", "y", "adad",
			})
		})

		Convey("test find all", func() {
			pattern := regexp.MustCompile("(a[bcd])(a[bcd])")
			So(pattern.FindAllStringSubmatch("abab x acac y adad", -1), ShouldResemble, [][]string{
				{"abab", "ab", "ab"},
				{"acac", "ac", "ac"},
				{"adad", "ad", "ad"},
			})
			So(pattern.FindAllString("abab x acac y adad", -1), ShouldResemble, []string{
				"abab", "acac", "adad",
			})
		})

		Convey("oss pattern", t, func() {
			pattern := regexp.MustCompile("(^oss://([^/]+)/([\\s\\S]*)$)|(^stream://$)")
			So(pattern.MatchString("oss://{bucket}/render/{barename}.{autoext}"), ShouldBeTrue)
		})
	})
}
