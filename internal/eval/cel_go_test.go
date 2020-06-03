package eval

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCelGo(t *testing.T) {
	Convey("TestCelGo", t, func() {
		Convey("example for github", func() {
			env, err := cel.NewEnv(cel.Declarations(
				decls.NewVar("name", decls.String),
				decls.NewVar("group", decls.String),
			))
			So(err, ShouldBeNil)
			ast, issues := env.Compile(`name.startsWith("/groups/" + group)`)
			So(issues, ShouldBeNil)
			prg, err := env.Program(ast)
			So(err, ShouldBeNil)
			val, details, err := prg.Eval(map[string]interface{}{
				"name":  "/groups/acme.co/documents/secret-stuff",
				"group": "acme.co",
			})
			So(err, ShouldBeNil)
			So(val.Value(), ShouldBeTrue)
			So(details, ShouldBeNil)
		})

		Convey("map[string]int", func() {
			env, err := cel.NewEnv(cel.Declarations(
				decls.NewVar("key1", decls.Int),
				decls.NewVar("key2", decls.Int),
			))
			So(err, ShouldBeNil)
			ast, issues := env.Compile(`key1 * (key2 + 1)`)
			So(issues, ShouldBeNil)
			prg, err := env.Program(ast)
			So(err, ShouldBeNil)

			val, details, err := prg.Eval(map[string]interface{}{
				"key1": 2,
				"key2": 4,
			})
			So(err, ShouldBeNil)
			So(val, ShouldEqual, 10)
			So(details, ShouldBeNil)
		})
	})
}
