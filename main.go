package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jhi721/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	s := &state{
		config: &cfg,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	if err := commands.register("login", handlerLogin); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println("not enough args were provided")
		os.Exit(1)
	}

	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	if err := commands.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
