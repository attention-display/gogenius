package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type iconst struct {
	items *Group
}

func Const() *iconst {
	i := &iconst{
		items: newGroup("(", ")", "\n"),
	}
	i.items.omitWrapIf = func() bool {
		// We only need to omit wrap while length == 1.
		// NewIf length == 0, we need to keep it, or it will be invalid expr.
		return i.items.length() == 1
	}
	return i
}
func (i *iconst) render(w io.Writer) {
	utils.WriteString(w, "const ")
	i.items.render(w)
}

func (i *iconst) AddField(name, value interface{}) *iconst {
	i.items.append(field(name, value, "="))
	return i
}
func (i *iconst) AddTypedField(name, typ, value interface{}) *iconst {
	i.items.append(typedField(name, typ, value, "="))
	return i
}

func (i *iconst) AddLineComment(content string, args ...interface{}) *iconst {
	i.items.append(LineComment(content, args...))
	return i
}
