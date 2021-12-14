package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/Scopics/architecture-lab-4/commands"
	"github.com/Scopics/architecture-lab-4/engine"
)

func main() {
	path := "cmd/example/example.txt"
	file := flag.String("f", path, "Example file")
	flag.Parse()

	input, err := os.Open(*file)
	if err != nil {
		log.Fatalf("error in open file %s", err)
		return
	}
	defer input.Close()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := commands.Parse(commandLine)
		eventLoop.Post(cmd)
	}

	eventLoop.AwaitFinish()
}
