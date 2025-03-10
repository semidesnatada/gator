package commands

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ Command) error {
	err := s.DB.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("successfully removed all users from database")
	return err
}