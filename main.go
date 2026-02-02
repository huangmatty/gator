package main

import (
	"log"
	"os"

	"github.com/huangmatty/gator/internal/config"
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
	cmds := commands{
		commandsToHandlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("usage: gator <command> [args...]")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
