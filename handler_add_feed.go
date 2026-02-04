package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		fmt.Printf("usage: %s <feed_name> <feed_url>\n", cmd.name)
		return nil
	}

	name, url := cmd.args[0], cmd.args[1]
	_, err := s.db.GetFeedByUrl(context.Background(), url)
	if err == nil {
		fmt.Printf("rss feed \"%s\" already exists\n", url)
		return nil
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("failed to add feed: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to add feed: %w", err)
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to add feed follow: %w", err)
	}

	fmt.Println("**New feed added and followed**")
	fmt.Printf("ID:      %v\n", feed.ID)
	fmt.Printf("Name:    %v\n", feed.Name)
	fmt.Printf("URL:     %v\n", feed.Url)
	fmt.Printf("Created: %v\n", feed.CreatedAt.UTC())
	return nil
}
