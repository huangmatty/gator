package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("users table reset unsuccessful: %w", err)
	}
	fmt.Println("successfully reset users table")
	if err := s.db.DeleteFeeds(context.Background()); err != nil {
		return fmt.Errorf("feeds table reset unsuccessful: %w", err)
	}
	fmt.Println("successfully reset feeds table")
	return nil
}
