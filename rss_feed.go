package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// create http client
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// creat GET request with context
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	// send GET request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request to %s: %w", feedURL, err)
	}
	defer resp.Body.Close()

	// check response status before parsing data
	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("bad response status: %v", resp.Status)
	}

	// read raw bytes from response body
	feedData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response from %s: %w", feedURL, err)
	}

	// unmarshal data to struct
	var feed RSSFeed
	if err := xml.Unmarshal(feedData, &feed); err != nil {
		return nil, fmt.Errorf("error unmarshalling data from %s: %w", feedURL, err)
	}
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}

	return &feed, nil
}
