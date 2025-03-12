package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/semidesnatada/gator/internal/database"
)

func handlerFollow(s *state, c Command) error {

	user_id, u_err := s.DB.GetUserID(context.Background(), s.Config.CurrentUserName)
	if u_err != nil {
		return u_err
	}
	
	feed_id, f_err := s.DB.GetFeedIDFromUrl(context.Background(), c.Argument[0])
	if f_err != nil {
		return f_err
	}


	newRecordDetails, dbErr := s.DB.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{ID: uuid.New(),
										CreatedAt: time.Now(),
										UpdatedAt: time.Now(),
										UserID: user_id,
										FeedID: feed_id,
									})
	if dbErr != nil {
		return dbErr
	}

	fmt.Printf("user: %s has successfully followed feed: %s\n", newRecordDetails.UserName, newRecordDetails.FeedName)

	return nil
}