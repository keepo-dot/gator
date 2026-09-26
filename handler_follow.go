package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/keepo-dot/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("follow command requires a url")
	}
	url := cmd.arguments[0]
	userID := user.ID
	feedID, err := s.db.GetFeedIDByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed ID: %w", err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}
	fmt.Printf("Feed followed. \nFeed Name: %s \nUser Name: %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}
