package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected one arg with the username")
	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Println("Username was set successfuly")

	return nil
}
