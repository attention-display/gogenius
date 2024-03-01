package gogenius

import (
	"testing"

	"github.com/attention-display/gogenius/utils"
)

func TestImports(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `import (
// test
"context"
. "time"
_ "math"

test "testing"
)
`
	Import().
		AddLineComment("test").
		AddPath("context").
		AddDot("time").
		AddBlank("math").
		AddLine().
		AddAlias("testing", "test").
		render(buf)

	utils.CompareAST(t, expected, buf.String())
}
