package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

type delete struct {
	Arg    string
	Symbol string
}

func (p *delete) Execute(loop engine.Handler) {
	res := strings.ReplaceAll(p.Arg, p.Symbol, "")
	loop.Post(&print{Arg: res})
}
