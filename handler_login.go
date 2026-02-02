package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}
	name := cmd.args[0]
	if err := s.config.SetUser(name); err != nil {
		return fmt.Errorf("unable to set current user: %w", err)
	}
	fmt.Printf("successfully set user -- %s\n", name)
	return nil
}
