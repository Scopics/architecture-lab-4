package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

type delete struct {
	arg    string
	symbol string
}

func (p *delete) Execute(loop engine.Handler) {
	res := strings.ReplaceAll(p.arg, p.symbol, "")
	loop.Post(&print{Arg: res})
}
