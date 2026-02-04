package main

import (
	"log"
	"os"

	"github.com/CybrRonin/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	//fmt.Printf("Read config: %v\n", cfg)

	currState := &state{&cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	commandArgs := os.Args

	if len(commandArgs) < 2 {
		log.Fatal("Not enough arugments to proceed")
	}

	cmdName := commandArgs[1]
	args := commandArgs[2:]

	cmd := command{Name: cmdName, Args: args}

	err = cmds.run(currState, cmd)
	if err != nil {
		log.Fatal(err)
	}
	/*
	   cfg, err = config.Read()

	   	if err != nil {
	   		log.Fatalf("error reading config: %v", err)
	   	}

	   fmt.Printf("Read config again: %v\n", cfg)
	*/
}
