package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

func Parse(line string) engine.Command {
	splittedLine := strings.Fields(line)

	command, params := splittedLine[0], splittedLine[1:]
	return Compose(command, params)
}
