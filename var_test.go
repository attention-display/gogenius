package gogenius

import (
	"testing"

	"github.com/attention-display/gogenius/utils"
)

func TestVar(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "var Version=2"

		Var().
			AddField("Version", Lit(2)).
			render(buf)

		utils.CompareAST(t, expected, buf.String())
	})

	t.Run("typed", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := "var Version int =2"

		Var().
			AddTypedField("Version", "int", Lit(2)).
			render(buf)

		utils.CompareAST(t, expected, buf.String())
	})

	t.Run("multiple", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `var (
Version=2
Description="Hello, World!"
)
`

		Var().
			AddField("Version", Lit(2)).
			AddField("Description", Lit("Hello, World!")).
			render(buf)

		utils.CompareAST(t, expected, buf.String())
	})

	t.Run("decl", func(t *testing.T) {
		buf := pool.Get()
		defer buf.Free()

		expected := `var _ io.Reader`

		Var().
			AddDecl("_", "io.Reader").
			render(buf)

		utils.CompareAST(t, expected, buf.String())
	})
}
