package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jhi721/gator/internal/database"
	"github.com/jhi721/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("please provide two args")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	userFromDb, err := s.queries.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.queries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    userFromDb.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
