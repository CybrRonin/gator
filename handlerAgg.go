package main

import (
	"context"
	"fmt"

	"github.com/CybrRonin/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching RSS feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
