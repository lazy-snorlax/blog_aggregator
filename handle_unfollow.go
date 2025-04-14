package main

import (
	"context"
	"fmt"

	"main.go/internal/database"
)

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	url := cmd.args[0]
	err := s.db.DeleteFeed(context.Background(), database.DeleteFeedParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Println("Successfully unfollowed feed")
	return nil
}
