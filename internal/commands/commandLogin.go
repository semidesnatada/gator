package commands

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, c Command) error {

	if len(c.Argument) == 0 {
		return errors.New("no username passed - login failed")
	} else if len(c.Argument) > 1 {
		return errors.New("too many arguments passed - login failed")
	}

	_, userExistsErr := s.DB.GetUser(context.Background(),c.Argument[0])

	if userExistsErr != nil {
		// return userExistsErr
		return errors.New("user not in database - cannot login")
	}

	loginError := s.Config.SetUser(c.Argument[0])
	if loginError != nil {
		return errors.New("couldn't login with that username - login failed")
	}
	fmt.Printf("user set to %s - login successful\n", c.Argument[0])
	return nil
}