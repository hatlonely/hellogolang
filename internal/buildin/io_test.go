package buildin

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCopy(t *testing.T) {
	Convey("test copy", t, func() {
		{
			reader := strings.NewReader("人生不发行往返车票，一旦出发了就再也不会归来了")
			writer := &bytes.Buffer{}
			n, err := io.Copy(writer, reader)
			So(err, ShouldBeNil)
			So(n, ShouldEqual, reader.Size())
			So(writer.String(), ShouldEqual, "人生不发行往返车票，一旦出发了就再也不会归来了")
		}
	})
}

func TestPipe(t *testing.T) {
	Convey("test pipe", t, func() {
		var wg sync.WaitGroup
		var buf []byte

		wg.Add(2)
		reader, writer := io.Pipe()
		go func() {
			_, _ = writer.Write([]byte("把活着的每一天看作生命的最后一天"))
			_ = writer.Close()
			wg.Done()
		}()
		go func() {
			buf, _ = ioutil.ReadAll(reader)
			wg.Done()
		}()
		wg.Wait()
		So(string(buf), ShouldEqual, "把活着的每一天看作生命的最后一天")
	})
}

func TestReader(t *testing.T) {
	Convey("test reader", t, func() {
		{
			reader := io.LimitReader(strings.NewReader("01234567890123456789"), 10)
			buf, _ := ioutil.ReadAll(reader)
			So(string(buf), ShouldEqual, "0123456789")
		}
		{
			reader := io.MultiReader(
				strings.NewReader("永远努力在你的生活之上保留一片天空\n"),
				strings.NewReader("成熟意味着停止展示自己，并学会隐藏自己\n"),
			)
			buf, _ := ioutil.ReadAll(reader)
			So(string(buf), ShouldEqual, "永远努力在你的生活之上保留一片天空\n成熟意味着停止展示自己，并学会隐藏自己\n")
		}
		{
			writer := &bytes.Buffer{}
			reader := io.TeeReader(
				strings.NewReader("生命的长短以时间来计算，生命的价值以贡献来计算"),
				writer,
			)
			buf, _ := ioutil.ReadAll(reader)
			So(string(buf), ShouldEqual, "生命的长短以时间来计算，生命的价值以贡献来计算")
			So(writer.String(), ShouldEqual, "生命的长短以时间来计算，生命的价值以贡献来计算")
		}
	})
}

func TestFileReadWrite(t *testing.T) {
	Convey("test file read/writer", t, func() {
		Convey("test ioutil", func() {
			So(ioutil.WriteFile("test.txt", []byte("人生天地之间，若白驹过隙，忽然而已"), 0644), ShouldBeNil)
			buf, _ := ioutil.ReadFile("test.txt")
			So(string(buf), ShouldEqual, "人生天地之间，若白驹过隙，忽然而已")
		})
		Convey("test bufio", func() {
			{
				fp, _ := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
				writer := bufio.NewWriter(fp)
				_, _ = writer.WriteString("居安思危\n思则有备\n有备无患")
				_ = writer.Flush()
				fp.Close()
			}
			{
				fp, _ := os.Open("test.txt")
				reader := bufio.NewReader(fp)
				buf := bytes.Buffer{}
				for {
					line, err := reader.ReadString('\n')
					buf.WriteString(line)
					if err != nil {
						break
					}
				}
				So(buf.String(), ShouldEqual, "居安思危\n思则有备\n有备无患")
				fp.Close()
			}
			{
				fp, _ := os.Open("test.txt")
				scanner := bufio.NewScanner(fp)
				buf := bytes.Buffer{}
				for scanner.Scan() {
					buf.WriteString(scanner.Text() + "\n")
				}
				So(buf.String(), ShouldEqual, "居安思危\n思则有备\n有备无患\n")
				fp.Close()
			}
		})
		Convey("test fmt", func() {
			{
				fp, _ := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
				_, _ = fmt.Fprint(fp, "居安思危\n思则有备\n有备无患")
				fp.Close()
			}
			{
				fp, _ := os.Open("test.txt")
				buf := bytes.Buffer{}
				for {
					var line string
					_, err := fmt.Fscan(fp, &line)
					if err != nil {
						break
					}
					buf.WriteString(line + "\n")
				}
				So(buf.String(), ShouldEqual, "居安思危\n思则有备\n有备无患\n")
			}
		})
		_ = os.RemoveAll("test.txt")
	})
}

func TestDirTravel(t *testing.T) {
	Convey("Give 遍历上个层级目录", t, func() {
		root := ".."
		Convey("When 使用filepath遍历（递归）", func() {
			var paths []string
			var names []string
			_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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
