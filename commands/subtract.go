package commands

import (
	"strconv"

	"github.com/Scopics/architecture-lab-4/engine"
)

type subtract struct {
	Arg1, Arg2 int
}

func (p *subtract) Execute(loop engine.Handler) {
	res := p.Arg1 - p.Arg2
	loop.Post(&print{Arg: strconv.Itoa(res)})
}
