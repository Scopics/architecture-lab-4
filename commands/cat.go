package commands

import (
	"github.com/Scopics/architecture-lab-4/engine"
)

type cat struct {
	Arg1, Arg2 string
}

func (p *cat) Execute(loop engine.Handler) {
	res := p.Arg1 + p.Arg2
	loop.Post(&print{Arg: res})
}
