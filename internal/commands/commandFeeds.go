package commands

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ Command) error {

	//function which gets a list of all the feeds in the db and 
	// for each, prints the name, user, url.

	feedsData, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("List of Feeds:")
	for i, item := range feedsData {
		fmt.Printf(" * feed %d: %s ; %s ; %s\n", i, item.Feedname, item.Username, item.Url)
	}

	return nil
}