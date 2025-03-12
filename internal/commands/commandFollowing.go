package commands

import (
	"context"
	"fmt"
)

func handlerFollowing (s *state, c Command) error {

	res, resErr := s.DB.GetFeedFollowsForUser(context.Background(), s.Config.CurrentUserName)
	if resErr != nil {
		return resErr
	}

	fmt.Printf("current user (%s) is following:\n", s.Config.CurrentUserName)
	for _, item := range res {
		fmt.Printf(" * %s\n",item.FeedName)
	}

	return nil
}