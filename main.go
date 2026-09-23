package main

import (
	"fmt"
	"os"

	"github.com/keepo-dot/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %s\n", err)
		os.Exit(1)
	}
	cliState := state{
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	userCmds := os.Args
	if len(userCmds) < 2 {
		fmt.Println("command name required")
		os.Exit(1)
	}
	cmd := command{
		name:      userCmds[1],
		arguments: userCmds[2:],
	}
	err = cmds.run(&cliState, cmd)
	if err != nil {
		fmt.Printf("error running %s command: %s\n", cmd.name, err)
		os.Exit(1)
	}
}
