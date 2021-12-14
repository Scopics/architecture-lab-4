package commands

import (
	"strings"

	"github.com/Scopics/architecture-lab-4/engine"
)

type Parser struct {
}

func (p *Parser) getCommand(commandName string, params []string) engine.Command {
	switch commandName {
	case "print":
		return &print{Arg: strings.Join(params, " ")}
	default:
		return &print{Arg: "Wrong Command Name"}
	}

}

func (p *Parser) parse(line string) engine.Command {
	splittedLine := strings.Fields(line)

	command, params := splittedLine[0], splittedLine[1:]
	return p.getCommand(command, params)
}
