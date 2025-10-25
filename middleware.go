package main

import (
	"context"

	"github.com/jhi721/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		userFromDb, err := s.queries.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return err
		}

		handler(s, c, userFromDb)

		return nil
	}
}
