package buildin

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOS(t *testing.T) {
	Convey("test os", t, func() {
		fmt.Println("group id:", os.Getgid())
		fmt.Println("user id:", os.Getuid())
		fmt.Println("effective user id:", os.Geteuid())
		fmt.Println("pid:", os.Getpid())
		fmt.Println("parent pid:", os.Getppid())
		fmt.Println("page size:", os.Getpagesize())

		workDir, _ := os.Getwd()
		fmt.Println("work directory:", workDir)

		hostname, _ := os.Hostname()
		fmt.Println("hostname:", hostname)

		cacheDir, _ := os.UserCacheDir()
		fmt.Println("cache directory:", cacheDir)
		configDir, _ := os.UserConfigDir()
		fmt.Println("config directory:", configDir)
		homeDir, _ := os.UserHomeDir()
		fmt.Println("home directory:", homeDir)
		tempDir := os.TempDir()
		fmt.Println("temp directory:", tempDir)

		executable, _ := os.Executable()
		fmt.Println("executable:", executable)
	})
}

func TestPath(t *testing.T) {
	Convey("test path", t, func() {
		So(os.RemoveAll("dir1"), ShouldBeNil)
		So(os.Mkdir("dir1", 0755), ShouldBeNil)
		So(os.Remove("dir1"), ShouldBeNil)                   // remove file or empty directory
		So(os.MkdirAll("dir1/dir2/dir3", 0755), ShouldBeNil) // mkdir -p
		So(os.RemoveAll("dir1"), ShouldBeNil)                // rm -rf

		fp, err := os.Create("test.txt")
		So(fp, ShouldNotBeNil)
		So(err, ShouldBeNil)
		fp.Close()
		So(os.Rename("test.txt", "test1.txt"), ShouldBeNil)
		So(os.RemoveAll("test1.txt"), ShouldBeNil)
	})

	Convey("test file path", t, func() {
		fmt.Println(filepath.Abs("."))

		So(filepath.Dir("/etc/nginx/conf.d/default.conf"), ShouldEqual, "/etc/nginx/conf.d")
		So(filepath.Base("/etc/nginx/conf.d/default.conf"), ShouldEqual, "default.conf")
		So(filepath.Ext("/etc/nginx/conf.d/default.conf"), ShouldEqual, ".conf")
		So(filepath.Clean("/../../ect/nginx//conf.d/.././../nginx/conf.d/default.conf"), ShouldEqual, "/ect/nginx/conf.d/default.conf")
		So(filepath.Join("/", "etc", "nginx/", "conf.d/default.conf"), ShouldEqual, "/etc/nginx/conf.d/default.conf")

		dir, filename := filepath.Split("/etc/nginx/conf.d/default.conf")
		So(dir, ShouldEqual, "/etc/nginx/conf.d/")
		So(filename, ShouldEqual, "default.conf")
	})
}

func TestFile(t *testing.T) {
	Convey("test file stat", t, func() {
		fp, err := os.Create("test.txt")
		So(err, ShouldBeNil)
		fmt.Println("fd:", fp.Fd())
		defer fp.Close()
		So(fp.Name(), ShouldEqual, "test.txt")
		info, err := fp.Stat()
		So(err, ShouldBeNil)
		fmt.Println(info.Name(), info.IsDir(), info.Size(), info.Mode(), info.ModTime())
		So(os.RemoveAll("test.txt"), ShouldBeNil)
	})

	Convey("test file read/write", t, func() {
		{
			fp, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			So(err, ShouldBeNil)
			_, err = fp.WriteString("优于别人并非高尚；今日之你优于昨日之你，才是真正的高尚")
			So(err, ShouldBeNil)
			_ = fp.Close()
		}
		{
			fp, err := os.OpenFile("test.txt", os.O_RDONLY, 0644)
			So(err, ShouldBeNil)
			buf := make([]byte, 128)
			n, err := fp.Read(buf)
			So(err, ShouldBeNil)
			So(string(buf[0:n]), ShouldEqual, "优于别人并非高尚；今日之你优于昨日之你，才是真正的高尚")
		}
		So(os.RemoveAll("test.txt"), ShouldBeNil)
	})
}
