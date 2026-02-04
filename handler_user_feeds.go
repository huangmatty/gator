package main

import (
	"context"
	"fmt"

	"github.com/huangmatty/gator/internal/database"
)

func handlerListUserFeeds(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to retrieve user's feeds: %w", err)
	}
	if len(feedFollows) < 1 {
		fmt.Printf("no feeds followed by %s\n", user.Name)
		return nil
	}
	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("Name: %v\n", feedFollow.FeedName)
	}
	return nil
}
