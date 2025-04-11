package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to reset users: %w", err)
	}
	fmt.Println("Successfully reset users")
	return nil
}
