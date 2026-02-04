package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		fmt.Printf("usage: %s <duration, e.g. 5s, 1m, 2h>\n", cmd.name)
		return nil
	}

	durationStr := cmd.args[0]
	time_between_reqs, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("failed to parse duration string: %w", err)
	}
	fmt.Printf("Collecting feeds every %v...\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Println("failed to scrape feeds:", err)
		}
	}
}
