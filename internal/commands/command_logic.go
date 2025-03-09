package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/semidesnatada/gator/internal/config"
	"github.com/semidesnatada/gator/internal/database"
)

type state struct {
	Config config.Config
	DB *database.Queries
}

type Command struct {
	Name string
	Argument []string
}

type commands struct {
	commandList map[string]func(*state, Command) error
}

func InitialiseState() (state, error) {
	con, err := config.Read()
	if err != nil {
		return state{}, err
	}
	return state{
		Config: con,
	}, nil
}

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


func GetCommands() commands {
	return commands{
		commandList: map[string]func(*state, Command) error{
			"login": handlerLogin,
			"register": handlerRegister,
		},
	}
}

func (c *commands) Run(s *state, cmd Command) error {
	
	cmdHandler, ok := c.commandList[cmd.Name]
	if !ok {
		return errors.New("command doesn't exist")
	}
	commandErr := cmdHandler(s, cmd)
	if commandErr != nil {
		return commandErr
	}	
	return nil
}