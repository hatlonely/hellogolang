package eval

import (
	"context"
	"testing"

	"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
	"github.com/antonmedv/expr"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func BenchmarkEval(b *testing.B) {
	b.Run("expr", func(b *testing.B) {
		prg, _ := expr.Compile(`a >= 10 && b < 20 && c == "abc"`)
		for i := 0; i < b.N; i++ {
			_, _ = expr.Run(prg, map[string]interface{}{
				"a": 11,
				"b": 19,
				"c": "abc",
			})
		}
	})

	b.Run("cel", func(b *testing.B) {
		env, _ := cel.NewEnv(cel.Declarations(
			decls.NewVar("a", decls.Int),
			decls.NewVar("b", decls.Int),
			decls.NewVar("c", decls.String),
		))
		ast, _ := env.Compile(`a >= 10 && b < 20 && c == "abc"`)
		prg, _ := env.Program(ast)
		for i := 0; i < b.N; i++ {
			_, _, _ = prg.Eval(map[string]interface{}{
				"a": 11,
				"b": 19,
				"c": "abc",
			})
		}
	})

	b.Run("gval", func(b *testing.B) {
		eval, _ := gval.Full().NewEvaluable(`a >= 10 && b < 20 && c == "abc"`)
		for i := 0; i < b.N; i++ {
			_, _ = eval(context.Background(), map[string]interface{}{
				"a": 11,
				"b": 19,
				"c": "abc",
			})
		}
	})

	b.Run("gval with json key", func(b *testing.B) {
		eval, _ := gval.Full(jsonpath.Language()).NewEvaluable(`a >= 10 && b < 20 && $["c"] == "abc"`)
		for i := 0; i < b.N; i++ {
			_, _ = eval(context.Background(), map[string]interface{}{
				"a": 11,
				"b": 19,
				"c": "abc",
			})
		}
	})

	b.Run("gval with string to int", func(b *testing.B) {
		eval, _ := gval.Full().NewEvaluable(`a >= 10 && b < 20 && c == "abc"`)
		for i := 0; i < b.N; i++ {
			_, _ = eval(context.Background(), map[string]interface{}{
				"a": "11",
				"b": "19",
				"c": "abc",
			})
		}
	})
}
