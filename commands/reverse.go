package commands

import (
	"github.com/Scopics/architecture-lab-4/engine"
)

type reverse struct {
	arg string
}

func (p *reverse) Execute(loop engine.Handler) {
	var res string
	for _, v := range p.arg {
		res = string(v) + res
	}
	loop.Post(&print{Arg: res})
}
