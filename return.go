package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type ireturn struct {
	items *Group
}

func Return(node ...interface{}) *ireturn {
	i := &ireturn{
		items: newGroup("", "", ", "),
	}
	i.items.append(node...)
	return i
}

func (i *ireturn) render(w io.Writer) {
	utils.WriteString(w, "return ")
	i.items.render(w)
}
