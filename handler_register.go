package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/keepo-dot/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("%s command requires a username", cmd.name)
	}
	name := cmd.arguments[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	s.cfg.SetUser(user.Name)
	fmt.Printf("user %s created:\n %#v", user.Name, user)
	return nil
}
