package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve feeds from database: %w", err)
	}
	if len(feeds) < 1 {
		fmt.Println("no feeds found")
		return nil
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			continue
		}
		fmt.Println("==============================")
		fmt.Printf("ID:      %v\n", feed.ID)
		fmt.Printf("Name:    %v\n", feed.Name)
		fmt.Printf("URL:     %v\n", feed.Url)
		fmt.Printf("Created: %v\n", feed.CreatedAt.UTC())
		fmt.Printf("User:    %v\n", user.Name)
	}
	return nil
}
