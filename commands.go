package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	commandsToHandlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commandsToHandlers[cmd.name]
	if !ok {
		return fmt.Errorf("%s handler not found", cmd.name)
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsToHandlers[name] = f
}
