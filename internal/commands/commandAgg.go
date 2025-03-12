package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/semidesnatada/gator/internal/database"
	"github.com/semidesnatada/gator/rss"
)

func handlerAgg(s *state, c Command) error {

	timeBetweenRequests := time.Second * 10
	ticker := time.NewTicker(timeBetweenRequests)

	fmt.Printf("Aggregating feeds every %vs.",timeBetweenRequests.Seconds())
	for ; ; <- ticker.C {
		scrapeErr := scrapeFeeds(s)
		if scrapeErr != nil {
			return scrapeErr
		}
	}

	// return nil
}

func scrapeFeeds(s *state) error {

	nextFeed, fetchErr := s.DB.GetNextFeedToFetch(context.Background())
	if fetchErr != nil {
		return fetchErr
	}

	markFetchedErr := s.DB.MarkFeedFetched(context.Background(),
	database.MarkFeedFetchedParams{ID: nextFeed.ID,
									LastFetched: sql.NullTime{Time: time.Now(), Valid:true},
									UpdatedAt: time.Now(),
	})
	if markFetchedErr != nil {
		return markFetchedErr
	}

	feedContent, reqErr := rss.FetchFeed(context.Background(), nextFeed.Url)
	if reqErr != nil {
		return reqErr
	}

	feedContent.FeedUnescape()

	storeErr := storeFeeds(s,feedContent, nextFeed.ID)
	if storeErr != nil {
		return storeErr
	}

	// printFeeds(feedContent)

	return nil
}

func storeFeeds(s *state, feed rss.RSSFeed, nextFeedID uuid.UUID) error {

	for _, item := range feed.Channel.Item {
		exists, err := s.DB.CheckIfPostAlreadyStored(context.Background(), item.Link)
		if err != nil {
			return err
		}
		if !exists {
			//store record
			const timeParser = "Mon, 02 Jan 2006 15:04:05 -0700"
			formatedPub, formatErr := time.Parse(timeParser, item.PubDate)
			if formatErr != nil {
				return formatErr
			}
			_, postErr := s.DB.CreatePost(context.Background(),
			database.CreatePostParams{
				ID: uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title: item.Title,
				Url: item.Link,
				Description: item.Description,
				PublishedAt: formatedPub,
				// PublishedAt: time.Now(),
				FeedID: nextFeedID,
			})
			if postErr != nil {
				return postErr
			}
		}
	}


	return nil
}

// func printFeeds(feed rss.RSSFeed) {
// 	for no, item := range feed.Channel.Item {
// 		fmt.Println("====================")
// 		fmt.Printf("Feed Item Number: %d\n", no+1)		
// 		fmt.Printf("Title: 			%s\n", item.Title)
// 		fmt.Printf("Link: 			%s\n", item.Link)
// 		// fmt.Printf("Description:	%s\n", item.Description)
// 		fmt.Printf("Date: 			%s\n", item.PubDate)
// 		fmt.Println()
// 	}	
// }