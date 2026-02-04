package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("failed to reset users table: %w", err)
	}
	fmt.Println("successfully reset users table")
	if err := s.db.DeleteFeeds(context.Background()); err != nil {
		return fmt.Errorf("failed to reset feeds table: %w", err)
	}
	fmt.Println("successfully reset feeds table")
	return nil
}
