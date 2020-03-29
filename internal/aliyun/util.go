package aliyun

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hpifu/go-kit/hconf"
)

func LoadOSSConfig() (string, string, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", "", err
	}

	conf, err := hconf.New("ini", "local", filepath.Join(home, ".ossutilconfig"))
	if err != nil {
		return "", "", "", err
	}

	fmt.Println(conf)

	return conf.GetDefaultString("Credentials.endpoint"),
		conf.GetDefaultString("Credentials.accessKeyID"),
		conf.GetDefaultString("Credentials.accessKeySecret"),
		nil
}
