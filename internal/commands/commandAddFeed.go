package commands

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/semidesnatada/gator/internal/database"
)

func handlerAddFeed(s *state, c Command) error {
	if len(c.Argument) < 2 {
		return errors.New("too few arguments - addFeed failed")
	} else if len(c.Argument) > 2 {
		return errors.New("too many arguments - addFeed failed")
	}

	currentUserName := s.Config.CurrentUserName

	currentUser, idErr := s.DB.GetUser(context.Background(),currentUserName)
	if idErr != nil {
		return errors.New("not logged in - can't add feed")
	}

	_, err := s.DB.CreateFeed(context.Background(),
	database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: c.Argument[0],
		Url: c.Argument[1],
		UserID: currentUser.ID,
	})
	if err != nil {
		return err
	}

	followErr := handlerFollow(s, Command{Name:"follow", Argument: []string{c.Argument[1]}})
	if followErr != nil {
		return followErr
	}

	return nil
}