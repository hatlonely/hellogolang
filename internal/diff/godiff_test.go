package diff

import (
	"fmt"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestGoDiff(t *testing.T) {
	const (
		text1 = "Lorem ipsum dolor."
		text2 = "Lorem dolor sit amet."
	)
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, true)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
