package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Printf("usage: %s <name>\n", cmd.name)
		return nil
	}

	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == sql.ErrNoRows {
		fmt.Printf("username \"%s\" doesn't exit\n", name)
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to log in: %w", err)
	}

	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}
	fmt.Printf("successfully set current user as \"%s\"\n", name)
	return nil
}
