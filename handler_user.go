package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jhi721/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected one arg with the username")
	}

	name := cmd.args[0]

	_, err := s.queries.GetUser(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.config.SetUser(name); err != nil {
		return err
	}

	fmt.Println("Username was set successfuly")

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected one arg with the username")
	}

	name := cmd.args[0]

	userFromDb, _ := s.queries.GetUser(context.Background(), name)
	if userFromDb.Name != "" {
		log.Fatal("User already exists")
	}

	user, err := s.queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("User was created")
	fmt.Println(user)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.queries.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		currentStatus := ""
		if user.Name == s.config.CurrentUserName {
			currentStatus = "(current)"
		}

		fmt.Printf("* %s %s\n", user.Name, currentStatus)
	}

	return nil
}
