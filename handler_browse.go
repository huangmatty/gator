package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/huangmatty/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Println("limit ust be a number")
			return nil
		}
		if parsedLimit <= 0 {
			fmt.Println("limit ust be a number greater than 0")
			return nil
		}
		limit = parsedLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve posts for user: %w", err)
	}

	for _, post := range posts {
		description := "no description"
		if post.Description.Valid {
			description = post.Description.String
		}
		fmt.Println("==============================")
		fmt.Printf("Published:   %v\n", post.PublishedAt.Time.UTC())
		fmt.Printf("Title:       %s\n", post.Title)
		fmt.Printf("URL:         %s\n", post.Url)
		fmt.Printf("Description: %s\n", description)
		fmt.Printf("FeedID:      %v\n", post.FeedID)
	}
	return nil
}
