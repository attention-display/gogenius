package gogenius

import (
	"io"

	"github.com/attention-display/gogenius/utils"
)

type iif struct {
	judge Node
	body  *Group
}

func If(judge interface{}) *iif {
	return &iif{
		judge: parseNode(judge),
		body:  newGroup("{\n", "\n}", "\n"),
	}
}
func (i *iif) render(w io.Writer) {
	utils.WriteString(w, "if ")
	i.judge.render(w)
	i.body.render(w)
}

func (i *iif) AddBody(node ...interface{}) *iif {
	i.body.append(node...)
	return i
}

type ifor struct {
	judge Node
	body  *Group
}

func (i *ifor) render(w io.Writer) {
	utils.WriteString(w, "for ")
	i.judge.render(w)
	i.body.render(w)
}

func For(judge interface{}) *ifor {
	return &ifor{
		judge: parseNode(judge),
		body:  newGroup("{\n", "\n}", "\n"),
	}
}

func (i *ifor) AddBody(node ...interface{}) *ifor {
	i.body.append(node...)
	return i
}

type icase struct {
	judge Node // judge == nil means it's a default case.
	body  *Group
}

func (i *icase) render(w io.Writer) {
	if i.judge == nil {
		utils.WriteString(w, "default:")
	} else {
		utils.WriteString(w, "case ")
		i.judge.render(w)
		utils.WriteString(w, ":")
	}
	i.body.render(w)
}

func (i *icase) AddBody(node ...interface{}) *icase {
	i.body.append(node...)
	return i
}

type iswitch struct {
	judge       Node
	cases       []*icase
	defaultCase *icase
}

func (i *iswitch) render(w io.Writer) {
	utils.WriteString(w, "switch ")
	i.judge.render(w)
	utils.WriteString(w, "{\n")
	for _, c := range i.cases {
		c.render(w)
		utils.WriteString(w, "\n")
	}
	if i.defaultCase != nil {
		i.defaultCase.render(w)
		utils.WriteString(w, "\n")
	}
	utils.WriteString(w, "}")
}

func Switch(judge interface{}) *iswitch {
	return &iswitch{
		judge: parseNode(judge),
	}
}
func (i *iswitch) NewCase(judge Node) *icase {
	ic := &icase{
		judge: judge,
		body:  newGroup("\n", "", "\n"),
	}
	i.cases = append(i.cases, ic)
	return ic
}

func (i *iswitch) NewDefault() *icase {
	ic := &icase{
		body: newGroup("\n", "", "\n"),
	}
	i.cases = append(i.cases, ic)
	return ic
}

func Continue() Node {
	return String("continue")
}
