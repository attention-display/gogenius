package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type ipackage struct {
	name string
}

func (i *ipackage) render(w io.Writer) {
	utils.WriteStringF(w, "package %s\n", i.name)
}

func Package(name string) *ipackage {
	return &ipackage{name: name}
}
