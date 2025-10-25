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

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("please provide two args")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.queries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.queries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return nil
	}

	fmt.Println(feed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.queries.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("please provide url as arg")
	}

	url := cmd.args[0]

	feedFromDb, err := s.queries.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	follow, err := s.queries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedFromDb.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(follow.FeedName)
	fmt.Println(follow.UserName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.queries.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}

	return nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("please provide url arg")
	}

	url := cmd.args[0]

	feedFromDb, err := s.queries.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.queries.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedFromDb.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
