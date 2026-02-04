package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err == sql.ErrNoRows {
		return fmt.Errorf("feed doesn't exist; use command addfeed to add feed: %w", err)
	}
	if err != nil {
		return fmt.Errorf("unable to follow feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to follow feed: %w", err)
	}
	fmt.Println("New feed followed:")
	fmt.Printf("ID:      %v\n", feedFollow.ID)
	fmt.Printf("Name:    %v\n", feedFollow.FeedName)
	fmt.Printf("User:    %v\n", feedFollow.UserName)
	fmt.Printf("Created: %v\n", feedFollow.CreatedAt.UTC())
	return nil
}
