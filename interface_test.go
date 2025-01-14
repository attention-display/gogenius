package gogenius

import (
	"testing"

	"github.com/attention-display/gogenius/utils"
)

func TestInterface(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := `type Tester interface {
// TestA is a test
TestA(a int64, b int)

TestB() (err error)
}
`
	in := Interface("Tester")
	in.AddLineComment("TestA is a test")
	in.NewFunction("TestA").
		AddParameter("a", "int64").
		AddParameter("b", "int")
	in.AddLine()
	in.NewFunction("TestB").
		AddResult("err", "error")

	in.render(buf)

	utils.CompareAST(t, expected, buf.String())
}
