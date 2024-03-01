package gogenius

import (
	"testing"

	"github.com/attention-display/gogenius/utils"
)

func TestDefer(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "defer hello()"

	Defer("hello()").render(buf)

	utils.CompareAST(t, expected, buf.String())
}
