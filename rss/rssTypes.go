package rss

import "html"

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (f *RSSFeed) FeedUnescape() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Link = html.UnescapeString(f.Channel.Link)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)
	for i, item := range f.Channel.Item {
		item.ItemUnescape()
		f.Channel.Item[i] = item
	}
}

func (i *RSSItem) ItemUnescape() {
	i.Title = html.UnescapeString(i.Title)
	i.Link = html.UnescapeString(i.Link)
	i.Description = html.UnescapeString(i.Description)
	i.PubDate = html.UnescapeString(i.PubDate)
}