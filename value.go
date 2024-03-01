package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type ivalue struct {
	typ   string
	items *Group
}

func Value(typ string) *ivalue {
	// FIXME: we need to support builtin type like map and slide.
	return &ivalue{
		typ:   typ,
		items: newGroup("{", "}", ","),
	}
}

func (v *ivalue) render(w io.Writer) {
	utils.WriteString(w, v.typ)
	v.items.render(w)
}

func (v *ivalue) String() string {
	buf := pool.Get()
	defer buf.Free()

	v.render(buf)
	return buf.String()
}

func (v *ivalue) AddField(name, value interface{}) *ivalue {
	v.items.append(field(name, value, ":"))
	return v
}
