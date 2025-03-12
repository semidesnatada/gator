package commands

import (
	"context"
	"fmt"

	"github.com/semidesnatada/gator/internal/database"
)

func handlerUnfollow(s *state, c Command) error {

	user_id, u_err := s.DB.GetUserID(context.Background(), s.Config.CurrentUserName)
	if u_err != nil {
		return u_err
	}

	feed_id, f_err := s.DB.GetFeedIDFromUrl(context.Background(), c.Argument[0])
	if f_err != nil {
		return f_err
	}

	deleteErr := s.DB.DeleteFeedFollow(context.Background(),
	database.DeleteFeedFollowParams{UserID: user_id,
									FeedID: feed_id,
								})
	if deleteErr != nil {
		return deleteErr
	}

	fmt.Printf("user: %s has successfully unfollowed feed: %s", s.Config.CurrentUserName, c.Argument[0])

	return nil
}