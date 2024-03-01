package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type idefer struct {
	body Node
}

func (i *idefer) render(w io.Writer) {
	utils.WriteString(w, "defer ")
	i.body.render(w)
}

func Defer(body interface{}) Node {
	// Add extra space here.
	return &idefer{parseNode(body)}
}
