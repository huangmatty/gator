package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Printf("usage: %s <name>\n", cmd.name)
		return nil
	}

	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Printf("username \"%s\" already exists\n", name)
		return nil
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("failed to create user: %w", err)
	}

	user, err := s.db.CreateUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}
	fmt.Printf("ID:      %v\n", user.ID)
	fmt.Printf("Name:    %v\n", user.Name)
	fmt.Printf("Created: %v\n", user.CreatedAt.UTC())
	return nil
}
