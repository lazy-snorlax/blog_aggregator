package main

import (
	"context"
	"fmt"
)

func handleFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("an error occured %w", err)
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("orphaned feed %v", feed.ID)
		} else {
			fmt.Printf("* %s   %s   %s\n", feed.Name, feed.Url, user.Name)
		}
	}
	return nil
}
