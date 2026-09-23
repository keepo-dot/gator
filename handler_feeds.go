package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds in db: %w", err)
	}
	if len(feeds) < 1 {
		return fmt.Errorf("no feeds in database")
	}
	for _, feed := range feeds {
		fmt.Printf("ID: %s\n", feed.ID)
		fmt.Printf("Feed Name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Created at: %s\n", feed.CreatedAt)
		if feed.CreatedAt != feed.UpdatedAt {
			fmt.Printf("Updated at: %s\n", feed.UpdatedAt)
		}
		fmt.Printf("Added by: %s\n", feed.UserName)
	}
	return nil
}
