package commands

import (
	"strconv"

	"github.com/Scopics/architecture-lab-4/engine"
)

type add struct {
	arg1, arg2 int
}

func (p *add) Execute(loop engine.Handler) {
	res := p.arg1 + p.arg2
	loop.Post(&print{Arg: strconv.Itoa(res)})
}
