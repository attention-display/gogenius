package utils

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func WriteString(w io.Writer, s ...string) {
	for _, v := range s {
		_, err := w.Write([]byte(v))
		if err != nil {
			panic(fmt.Errorf("write string: %v", err))
		}
	}
}

func WriteStringF(w io.Writer, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	_, err := w.Write([]byte(s))
	if err != nil {
		panic(fmt.Errorf("write string: %v", err))
	}
}

// cleanAST will remove all space and newline from code.
// We know it will break the AST semantic, but golang doesn't support parse
// partial code source, we have to do like this.
// Maybe we can find a better way to compare the AST in go.
func CleanAST(a string) string {
	a = strings.ReplaceAll(a, " ", "")
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, "\t", "")
	return a
}

func CompareAST(t *testing.T, a, b string) {
	na, nb := CleanAST(a), CleanAST(b)
	if na == nb {
		return
	}
	t.Error("AST is not the same.")
	t.Errorf("left:\n%s\ncleaned:\n%s", a, na)
	t.Errorf("right:\n%s\ncleaned:\n%s", b, nb)
}
