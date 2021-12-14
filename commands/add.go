package commands

import (
	"strconv"

	"github.com/Scopics/architecture-lab-4/engine"
)

type add struct {
	Arg1, Arg2 int
}

func (p *add) Execute(loop engine.Handler) {
	res := p.Arg1 + p.Arg2
	loop.Post(&print{Arg: strconv.Itoa(res)})
}
