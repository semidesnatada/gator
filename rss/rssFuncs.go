package rss

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (RSSFeed, error) {

	req, reqErr := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if reqErr != nil {
		return RSSFeed{}, reqErr
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	res, resErr := client.Do(req)
	if resErr != nil {
		return RSSFeed{}, resErr
	}

	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)

	var output RSSFeed

	decodeErr := decoder.Decode(&output)
	if decodeErr != nil {
		return RSSFeed{}, decodeErr
	}

	output.Channel.Title = html.UnescapeString(output.Channel.Title)

	return output, nil
}
