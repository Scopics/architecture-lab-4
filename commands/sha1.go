package commands

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/Scopics/architecture-lab-4/engine"
)

type sha struct {
	Arg string
}

func (p *sha) Execute(loop engine.Handler) {
	h := sha1.New()
	io.WriteString(h, p.Arg)
	res := fmt.Sprintf("% x", h.Sum(nil))
	loop.Post(&print{Arg: res})
}
