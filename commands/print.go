package commands

import (
	"fmt"

	"github.com/Scopics/architecture-lab-4/engine"
)

type print struct {
	Arg string
}

func (p *print) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}
