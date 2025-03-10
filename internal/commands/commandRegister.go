package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/semidesnatada/gator/internal/database"
)

func handlerRegister(s *state, c Command) error {
	if len(c.Argument) == 0 {
		return errors.New("no name passed - register failed")
	} else if len(c.Argument) > 1 {
		return errors.New("too many arguments passed - register failed")
	}

	_, userExistsErr := s.DB.GetUser(context.Background(),c.Argument[0])

	if userExistsErr == nil {
		// return userExistsErr
		return errors.New("user already in database - cannot register again")
	}

	_, registerError := s.DB.CreateUser(context.Background(), 
										database.CreateUserParams{
											ID: uuid.New(),
											CreatedAt: time.Now(),
											UpdatedAt: time.Now(),
											Name: c.Argument[0],
										})

	if registerError != nil {
		return registerError
	}

	fmt.Printf("user: %s, successfully registered\n", c.Argument[0])

	loginError := s.Config.SetUser(c.Argument[0])
	if loginError != nil {
		return errors.New("couldn't login as newly registered user - unsure why")
	}

	fmt.Printf("now logged in as: %s \n", c.Argument[0])

	return nil
}