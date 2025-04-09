package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("you must provide a username %s <name>", cmd.name)
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User set to %v \n", username)

	return nil
}
