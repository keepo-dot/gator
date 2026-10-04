package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/keepo-dot/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limitINT int
	var err error
	if len(cmd.arguments) == 0 {
		limitINT = 2
	} else {
		limitINT, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("error during conversion from string to int: %w", err)
		}
	}
	limitInt32 := int32(limitINT)
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limitInt32,
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Published At: %s\n", post.PublishedAt.Time)
		fmt.Printf("Created At: %s\n", post.CreatedAt.Format(time.RFC1123))
		if post.CreatedAt != post.UpdatedAt {
			fmt.Printf("Updated At: %s\n", post.UpdatedAt.Format(time.RFC1123))
		}
		fmt.Printf("URL: %s\n", post.Url)
	}
	return nil
}
