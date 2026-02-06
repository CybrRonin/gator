package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving users: %w", err)
	}
	for _, user := range users {
		name := user.Name
		response := fmt.Sprintf("* %s", name)
		if name == s.cfg.CurrentUserName {
			response = fmt.Sprintf("%s (current)", response)
		}
		fmt.Println(response)
	}
	return nil
}
