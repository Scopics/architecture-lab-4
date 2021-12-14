package commands

import (
	"github.com/Scopics/architecture-lab-4/engine"
)

type cat struct {
	arg1, arg2 string
}

func (p *cat) Execute(loop engine.Handler) {
	res := p.arg1 + p.arg2
	loop.Post(&print{Arg: res})
}
