package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.name)
	}

	name, url := cmd.args[0], cmd.args[1]
	_, err := s.db.GetFeedByUrl(context.Background(), url)
	if err == nil {
		return fmt.Errorf("rss feed \"%s\" already exists", url)
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("unable to add feed: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to add feed: %w", err)
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to add feed follow: %w", err)
	}

	fmt.Println("**New feed added and followed**")
	fmt.Printf("ID:      %v\n", feed.ID)
	fmt.Printf("Name:    %v\n", feed.Name)
	fmt.Printf("URL:     %v\n", feed.Url)
	fmt.Printf("Created: %v\n", feed.CreatedAt.UTC())
	return nil
}
