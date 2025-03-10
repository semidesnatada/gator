package commands

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, c Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	loggedIn := s.Config.CurrentUserName

	for _, user := range users {
		if user != loggedIn {
		fmt.Printf(" * %s\n", user)
	} else {
		fmt.Printf(" * %s (current)\n", user)
	}
	}

	return nil
}