package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		fmt.Printf("usage: %s <feed_url>\n", cmd.name)
		return nil
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err == sql.ErrNoRows {
		fmt.Println("feed doesn't exist; use command addfeed <feed_url> to add feed")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}
	fmt.Println("New feed followed:")
	fmt.Printf("ID:      %v\n", feedFollow.ID)
	fmt.Printf("Name:    %v\n", feedFollow.FeedName)
	fmt.Printf("User:    %v\n", feedFollow.UserName)
	fmt.Printf("Created: %v\n", feedFollow.CreatedAt.UTC())
	return nil
}
