package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	RSSFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching rss feed: %w", err)
	}
	fmt.Print(RSSFeed)
	return nil
}
