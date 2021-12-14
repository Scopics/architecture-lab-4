package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

type split struct {
	Arg string
	Sep string
}

func (p *split) Execute(loop engine.Handler) {
	res := strings.Split(p.Arg, p.Sep)
	for _, v := range res {
		loop.Post(&print{Arg: v})
	}
}
