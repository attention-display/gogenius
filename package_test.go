package gogenius

import (
	"testing"

	"github.com/attention-display/gogenius/utils"
)

func TestPackage(t *testing.T) {
	buf := pool.Get()
	defer buf.Free()

	expected := "package test\n"

	Package("test").render(buf)

	utils.CompareAST(t, expected, buf.String())
}
