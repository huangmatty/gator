package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huangmatty/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("username \"%s\" already exist", name)
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("unable to create user: %w", err)
	}

	user, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("unable to create user: %w", err)
	}
	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("unable to set current user: %w", err)
	}
	fmt.Printf("id: %v, created: %v, updated %v, name: %v\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}
