package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return fmt.Errorf("too many arguments for reset")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting table: %w", err)
	}
	fmt.Println("successfully reset table")
	return nil
}
