package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return errors.New("login handler expects a single username as an argument")
	}
	username := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error on user login: %w", err)
	}
	fmt.Printf("user has been set as: %s\n", username)
	return nil
}
