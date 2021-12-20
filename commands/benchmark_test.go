package commands

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	for i := 1; i <= 20; i++ {
		b.Run(fmt.Sprintf("%d-length", i*1000), func(b *testing.B) {
			Parse("palindrom " + strings.Repeat("A", i*10000))
		})
	}
}
