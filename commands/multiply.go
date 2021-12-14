package commands

import (
	"strconv"

	"github.com/Scopics/architecture-lab-4/engine"
)

type multiply struct {
	Arg1, Arg2 int
}

func (p *multiply) Execute(loop engine.Handler) {
	res := p.Arg1 * p.Arg2
	loop.Post(&print{Arg: strconv.Itoa(res)})
}
