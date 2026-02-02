package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == sql.ErrNoRows {
		return fmt.Errorf("username \"%s\" doesn't exit", name)
	}
	if err != nil {
		return fmt.Errorf("unable to login: %w", err)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("unable to set current user: %w", err)
	}
	fmt.Printf("successfully set user -- %s\n", name)
	return nil
}
