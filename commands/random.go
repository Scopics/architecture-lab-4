package commands

import (
	"math/rand"
	"strconv"

	"github.com/Scopics/architecture-lab-4/engine"
)

type random struct {
}

func (p *random) Execute(loop engine.Handler) {
	r := rand.New(rand.NewSource(99))
	res := r.Int()
	loop.Post(&print{Arg: strconv.Itoa(res)})
}
