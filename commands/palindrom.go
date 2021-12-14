package commands

import (
	"github.com/Scopics/architecture-lab-4/engine"
)

type palindrom struct {
	Arg string
}

func (p *palindrom) Execute(loop engine.Handler) {
	var res string
	for _, v := range p.Arg {
		res = string(v) + res
	}
	res = p.Arg + res
	loop.Post(&print{Arg: res})
}
