package buildin

import (
	"bufio"
	"crypto/subtle"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileRW(t *testing.T) {
	Convey("Given 给一个文件名以及文件的内容", t, func() {
		filename := "text.txt"
		context := "hello world!"

		Convey("When 使用ioutil写入文件内容", func() {
			err := ioutil.WriteFile(filename, []byte(context), 0644)
			So(err, ShouldBeNil)
		})

		Convey("When 使用ioutil读文件内容", func() {
			buf, err := ioutil.ReadFile(filename)
			So(err, ShouldBeNil)
			So(subtle.ConstantTimeCompare(buf, []byte(context)), ShouldEqual, 1)
		})

		Convey("When 使用os写入文件内容", func() {
			fp, err := os.Create(filename)
			defer fp.Close()
			So(err, ShouldBeNil)

			num, err := fp.WriteString(context)
			So(num, ShouldEqual, len(context))
			So(err, ShouldBeNil)

			err = fp.Sync()
			So(err, ShouldBeNil)
		})

		Convey("When 使用os读文件内容", func() {
			fp, err := os.Open(filename)
			defer fp.Close()
			So(err, ShouldBeNil)
			So(fp, ShouldNotBeNil)

			buf := make([]byte, 1024)
			num, err := fp.Read(buf)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, len(context))
			So(string(buf[0:num]), ShouldEqual, context)
		})

		Convey("When 使用bufio写入文件内容", func() {
			fp, err := os.Create(filename)
			defer fp.Close()
			So(err, ShouldBeNil)

			writer := bufio.NewWriter(fp)
			num, err := writer.WriteString(context)
			So(num, ShouldEqual, len(context))
			So(err, ShouldBeNil)

			err = writer.Flush()
			So(err, ShouldBeNil)
		})

		Convey("When 使用bufio读取文件内容", func() {
			fp, err := os.Open(filename)
			defer fp.Close()
			So(err, ShouldBeNil)

			reader := bufio.NewReader(fp)
			buf := make([]byte, 1024)
			num, err := reader.Read(buf)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, len(context))
			So(string(buf[0:num]), ShouldEqual, context)
		})

		Convey("When 使用fmt写入文件内容", func() {
			fp, err := os.Create(filename)
			defer fp.Close()
			So(err, ShouldBeNil)

			num, err := fmt.Fprintf(fp, context)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, len(context))
		})

		Convey("Finally 文件删除", func() {
			os.Remove(filename)
		})
	})
}

func TestFileTravel(t *testing.T) {
	Convey("Given 一个文件", t, func() {
		filename := "text.txt"
		context := "hello world!\nhello golang!\n"
		err := ioutil.WriteFile(filename, []byte(context), 0644)
		So(err, ShouldBeNil)

		Convey("When 使用bufio.scanner逐行读取", func() {
			fp, err := os.Open(filename)
			defer fp.Close()
			So(err, ShouldBeNil)

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				Println(scanner.Text())
			}

			So(scanner.Err(), ShouldBeNil)
		})

		Convey("When 使用bufio.reader逐行读取", func() {
			fp, err := os.Open(filename)
			defer fp.Close()
			So(err, ShouldBeNil)

			reader := bufio.NewReader(fp)
			var line string
			for {
				line, err = reader.ReadString('\n')
				if err != nil {
					break
				}
				Println(line[0 : len(line)-1])
			}
			So(err, ShouldEqual, io.EOF)
		})

		Convey("Finally 文件删除", func() {
			os.Remove(filename)
		})
	})
}

func TestDirTravel(t *testing.T) {
	Convey("Give 遍历上个层级目录", t, func() {
		root := ".."
		Convey("When 使用filepath遍历（递归）", func() {
			var paths []string
			var names []string
			filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				So(err, ShouldBeNil)
				paths = append(paths, path)
				names = append(names, info.Name())
				return nil
			})

			Convey("Then 结果中必须包含当前这个文件", func() {
				So(paths, ShouldContain, "../buildin/io_test.go")
				So(names, ShouldContain, "io_test.go")
			})
		})

		Convey("When 使用ioutil遍历（非递归）", func() {
			var files []string
			var dirs []string
			infos, err := ioutil.ReadDir(root)
			So(err, ShouldBeNil)
			for _, info := range infos {
				if info.IsDir() {
					dirs = append(dirs, info.Name())
				} else {
					files = append(files, info.Name())
				}
			}
			Convey("Then 结果中必须包含当前这个文件", func() {
				So(dirs, ShouldContain, "buildin")
				So(files, ShouldBeEmpty)
			})
		})
	})
}
