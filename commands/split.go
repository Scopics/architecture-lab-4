package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

type split struct {
	arg string
	sep string
}

func (p *split) Execute(loop engine.Handler) {
	res := strings.Split(p.arg, p.sep)
	for _, v := range res {
		loop.Post(&print{Arg: v})
	}
}
