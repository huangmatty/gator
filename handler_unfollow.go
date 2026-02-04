package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		fmt.Printf("usage: %s <feed_url>\n", cmd.name)
		return nil
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err == sql.ErrNoRows {
		fmt.Println("feed doesn't exist")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	_, err = s.db.GetFeedFollowForUser(context.Background(), database.GetFeedFollowForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err == sql.ErrNoRows {
		fmt.Printf("%s is not followed\n", url)
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	if err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}
	fmt.Printf("successfully unfollowed %s at %s\n", feed.Name, feed.Url)
	return nil
}
