package main

import (
	"context"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("unable to find current user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
