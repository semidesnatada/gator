package commands

import (
	"context"
	"fmt"

	"github.com/semidesnatada/gator/internal/database"
)


func handlerBrowse(s *state, c Command) error {

	user_id, u_err := s.DB.GetUserID(context.Background(), s.Config.CurrentUserName)
	if u_err != nil {
		return u_err
	}

	fetchedPosts, fetchErr := s.DB.GetPostsForUser(context.Background(),
	database.GetPostsForUserParams{
		UserID: user_id,
		Limit: 5,
	})
	if fetchErr != nil {
		return fetchErr
	}

	printPosts(fetchedPosts)

	return nil
}

func printPosts(posts []database.GetPostsForUserRow) {

	for _, post := range posts {
		fmt.Println("====================")
		fmt.Printf("Title: 			%s\n", post.Title)
		fmt.Printf("Link: 			%s\n", post.Url)
		// fmt.Printf("Description:	%s\n", post.Description)
		fmt.Printf("Date: 			%v\n", post.PublishedAt)
		fmt.Println()
	}

}
