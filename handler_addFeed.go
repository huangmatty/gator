package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/huangmatty/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: %s <feed name> <feed url>", cmd.name)
	}

	name, url := cmd.args[0], cmd.args[1]
	_, err := s.db.GetFeed(context.Background(), url)
	if err == nil {
		return fmt.Errorf("rss feed \"%s\" already exists", url)
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("unable to add feed: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("unable to add feed: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to add feed: %w", err)
	}
	fmt.Printf("ID:      %v\n", feed.ID)
	fmt.Printf("Name:    %v\n", feed.Name)
	fmt.Printf("URL:     %v\n", feed.Url)
	fmt.Printf("Created: %v\n", feed.CreatedAt.UTC())
	return nil
}
