package commands

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Scopics/architecture-lab-4/engine"
)

var cmd engine.Command

func BenchmarkParse(b *testing.B) {
	for i := 1; i <= 20; i++ {
		b.Run(fmt.Sprintf("%d-length", i*1000), func(b *testing.B) {
			cmd = Parse("palindrom " + strings.Repeat("A", i*10000))
		})
	}
}
