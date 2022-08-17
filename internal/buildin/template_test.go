package buildin

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTemplate(t *testing.T) {
	Convey("text/template", t, func() {
		tpl, err := template.New("").Funcs(template.FuncMap{
			"indent": func(indent int, v interface{}) string {
				if v == nil {
					return ""
				}
				str := v.(string)
				vals := strings.Split(str, "\n")
				var buf bytes.Buffer
				for _, val := range vals {
					for i := 0; i < indent; i++ {
						buf.WriteByte(' ')
					}
					buf.WriteString(val)
					buf.WriteByte('\n')
				}
				return buf.String()
			},
		}).Parse(`
key1: |
{{.Key1 | indent 4}}
{{.Key2}}
{{.Key3 | indent 4}}
`)
		So(err, ShouldBeNil)
		var buf bytes.Buffer
		So(tpl.Execute(&buf, map[string]interface{}{
			"Key1": "Hello\nWorld",
		}), ShouldBeNil)

		fmt.Println(buf.String())
	})
}

func TestEscape(t *testing.T) {
	Convey("TestEscape", t, func() {
		tpl, err := template.New("").Parse(`{} {{ "hello world" }}`)
		So(err, ShouldBeNil)
		var buf bytes.Buffer
		tpl.Execute(&buf, nil)

		fmt.Println(buf.String())
		So(buf.String(), ShouldEqual, "{} hello world")
	})
}
