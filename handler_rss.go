package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error forming http request: %w", err)
	}
	req.Header.Set("user-agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error during http response: %w", err)
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("http response returned with status code: %d", res.StatusCode)
	}
	defer res.Body.Close()
	xmlData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	var rssFeed RSSFeed
	if err = xml.Unmarshal(xmlData, &rssFeed); err != nil {
		return nil, fmt.Errorf("error during xml unmarshal: %w", err)
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Link = html.UnescapeString(rssFeed.Channel.Link)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Link = html.UnescapeString(rssFeed.Channel.Item[i].Link)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
		rssFeed.Channel.Item[i].PubDate = html.UnescapeString(rssFeed.Channel.Item[i].PubDate)
	}
	return &rssFeed, nil
}
