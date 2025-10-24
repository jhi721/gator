package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	if err := s.queries.DeleteUsers(context.Background()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Reset was successful!")

	return nil
}
