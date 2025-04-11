package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"main.go/internal/database"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	name := cmd.args[0]
	url := cmd.args[1]

	user_feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	fmt.Printf("%+v", user_feed)

	return nil
}
