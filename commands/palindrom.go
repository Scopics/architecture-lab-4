package commands

import (
	"github.com/Scopics/architecture-lab-4/engine"
)

type palindrom struct {
	arg string
}

func (p *palindrom) Execute(loop engine.Handler) {
	var res string
	for _, v := range p.arg {
		res = string(v) + res
	}
	res = p.arg + res

	loop.Post(&print{Arg: res})
}
