package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type itype struct {
	name string
	item Node
	sep  string
}

func Type(name string, typ interface{}) *itype {
	return &itype{
		name: name,
		item: parseNode(typ),
	}
}

func TypeAlias(name string, typ interface{}) *itype {
	return &itype{
		name: name,
		item: parseNode(typ),
		sep:  "=",
	}
}

func (i *itype) render(w io.Writer) {
	utils.WriteStringF(w, "type %s %s", i.name, i.sep)
	i.item.render(w)
}
