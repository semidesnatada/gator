package commands

import (
	"errors"

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

func GetCommands() commands {
	return commands{
		commandList: map[string]func(*state, Command) error{
			"login": handlerLogin,
			"register": handlerRegister,
			"reset": handlerReset,
			"users": handlerUsers,
			"agg": handlerAgg,
			"addfeed": handlerAddFeed,
			"feeds": handlerFeeds,
			"follow": handlerFollow,
			"following": handlerFollowing,
			"unfollow": handlerUnfollow,
			"browse": handlerBrowse,
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


// func requireLogin(handler func(s *state, c Command, u database.User) error) func(*state, Command) error {


// }