package reddit

import (
	"encoding/json"
	"fmt"
)

var (
	redditURL = "https://www.reddit.com"
)

type Awarding struct {
	AwardType   string `json:"award_type"`
	CoinPrice   int    `json:"coin_price"`
	Description string `json:"desctiption"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	IsEnabled   bool   `json:"is_enabled"`
}

type Data struct {
	After    string  `json:"after,omitempty"`
	Before   string  `json:"before,omitempty"`
	Children []*Link `json:"children,omitempty"`
	Dist     int     `json:"dist,omitempty"`
	Modhash  string  `json:"modhash,omitempty"`
}

type LinkFlairRichText struct {
	E string `json:"e"`
	T string `json:"t"`
}

type RedditCommentResponse []Data

type ListingArray []RedditListing

type RedditListing struct {
	Kind string
	Data Data `json:"data"`
}

func (r RedditCommentResponse) AllChildren() (t []*Link) {
	totalChildren := 0
	for _, v := range r {
		totalChildren += len(v.Children)
	}

	t = make([]*Link, 0, totalChildren)

	for _, listing := range r {
		t = append(t, listing.Children...)
	}
	return
}

func (r ListingArray) AllChildren() (t []*Link) {
	totalChildren := 0
	for _, v := range r {
		totalChildren += len(v.Data.Children)
	}

	t = make([]*Link, 0, totalChildren)

	for _, listing := range r {
		t = append(t, listing.Data.Children...)
	}
	return
}

func GetCommentListing(board, id, opts string) (r ListingArray, err error) {
	b, _, er := BrowserRequest(fmt.Sprintf(BoardUri(board)+"/comments/%s.json", id) + opts)
	if er != nil {
		err = er
		return
	}

	err = json.Unmarshal(b, &r)

	return
}

func getSub(boardName string) (*Subreddit, error) {
	var board = new(Subreddit)
	b, _, err := BrowserRequest(redditURL + fmt.Sprintf("/r/%s.json", boardName))

	if err != nil {
		return board, err
	}

	err = json.Unmarshal(b, board)
	if err != nil {
		return board, err
	}

	return board, nil
}
