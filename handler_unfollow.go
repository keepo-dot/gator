package main

import (
	"context"
	"fmt"

	"github.com/keepo-dot/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	userID := user.ID
	url := cmd.arguments[0]
	feedID, err := s.db.GetFeedIDByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed ID: %w", err)
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: userID,
		FeedID: feedID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed: %w", err)
	}
	return nil
}
