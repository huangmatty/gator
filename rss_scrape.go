package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/huangmatty/gator/internal/database"
	"github.com/lib/pq"
)

var timeLayouts = []string{
	time.RFC1123Z,
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	time.RFC3339,
}

func parsePubDate(pubDate string) sql.NullTime {
	for _, layout := range timeLayouts {
		if t, err := time.Parse(layout, pubDate); err == nil {
			return sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
	}
	return sql.NullTime{}
}

func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505"
	}
	return false
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to identify next feed to fetch: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}
	// fmt.Printf("Channel Title:       %s\n", rssFeed.Channel.Title)
	// fmt.Printf("Channel Link:        %s\n", rssFeed.Channel.Link)
	// fmt.Printf("Channel Description: %s\n", rssFeed.Channel.Description)
	// fmt.Println("Channel Items:")
	for _, item := range rssFeed.Channel.Item {
		publishedAt := parsePubDate(item.PubDate)
		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			Title: strings.TrimSpace(item.Title),
			Url:   item.Link,
			Description: sql.NullString{
				String: strings.TrimSpace(item.Description),
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if isUniqueViolation(err) {
				fmt.Println("skipping existing post")
				continue
			}
			log.Printf("failed to create post titled %s: %v", post.Title, err)
			continue
		}
		fmt.Printf("post \"%s\" added (%s)\n", post.Title, post.Url)
	}

	if err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID); err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}
	return nil
}
