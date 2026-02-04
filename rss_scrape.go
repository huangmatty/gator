package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to identify next feed to fetch: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	fmt.Printf("Channel Title:       %s\n", rssFeed.Channel.Title)
	fmt.Printf("Channel Link:        %s\n", rssFeed.Channel.Link)
	fmt.Printf("Channel Description: %s\n", rssFeed.Channel.Description)
	fmt.Println("Channel Items:")
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf(" - %s (published %v)\n", item.Title, item.PubDate)
	}

	if err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}
	return nil
}
