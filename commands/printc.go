package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

type printc struct {
	c   int
	arg string
}

func (p *printc) Execute(loop engine.Handler) {
	res := strings.Repeat(p.arg, p.c)
	loop.Post(&print{Arg: res})
}
