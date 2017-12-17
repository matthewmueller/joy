package testutil

import (
	"bytes"
	"strings"

	"github.com/sanity-io/litter"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// Format fn
func Format(expected, actual interface{}) string {
	e := litter.Sdump(expected)
	a := litter.Sdump(actual)

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(e, a, false)

	var buf bytes.Buffer
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			buf.WriteString("\x1b[102m\x1b[30m")
			buf.WriteString(diff.Text)
			buf.WriteString("\x1b[0m")
		case diffmatchpatch.DiffDelete:
			buf.WriteString("\x1b[101m\x1b[30m")
			buf.WriteString(diff.Text)
			buf.WriteString("\x1b[0m")
		case diffmatchpatch.DiffEqual:
			buf.WriteString(diff.Text)
		}
	}

	result := buf.String()
	result = strings.Replace(result, "\\n", "\n", -1)
	result = strings.Replace(result, "\\t", "\t", -1)
	return result
}
