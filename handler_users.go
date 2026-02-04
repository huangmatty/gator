package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve users from database: %w", err)
	}
	if len(users) < 1 {
		fmt.Println("no users found")
		return nil
	}
	currentUser := s.cfg.CurrentUsername
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Println("*", user.Name, "(current)")
			continue
		}
		fmt.Println("*", user.Name)
	}
	return nil
}
