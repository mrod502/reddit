package reddit

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	redditURL = "https://www.reddit.com"
)

type T3Data struct {
	ID                string              `json:"id"`
	Created           int64               `json:"created"`
	CreatedUTC        int64               `json:"created_utc"`
	LinkFlairRichText []LinkFlairRichText `json:"link_flair_richtext"`
	Author            string              `json:"author"`
	Title             string              `json:"title"`
	Selftext          string              `json:"selftext"`
	Ups               int                 `json:"ups"`
	Downs             int                 `json:"downs"`
	UpvoteRatio       float64             `json:"upvote_ratio"`
	AllAwardings      []Awarding          `json:"all_awardings"`
	Body              string              `json:"body"`
	Subreddit         string              `json:"subreddit"`
	Symbols           []string
}

func (t T3Data) Compact() CompactT3 {

	return CompactT3{
		Title:   t.Title,
		Text:    t.Selftext,
		Symbols: GetStockSymbols(t.Title + "\n" + t.Selftext),
		Created: time.Unix(t.Created, 0),
		Ups:     t.Ups,
	}
}

func (t T3Data) GetStockSymbols() T3Data {
	t.Symbols = GetStockSymbols(t.Selftext + "\n" + t.Title + "\n" + t.Body)
	return t
}

type Awarding struct {
	AwardType   string `json:"award_type"`
	CoinPrice   int    `json:"coin_price"`
	Description string `json:"desctiption"`
	Count       int    `json:"count"`
	Name        string `json:"name"`
	IsEnabled   bool   `json:"is_enabled"`
}

type Link struct {
	Kind string
	Data T3Data `json:"data"`
}

func (l Link) GetHrefs() (o []string) {
	matches := mdLinkRegex.FindAllStringSubmatch(l.Data.Selftext, -1)
	for _, m := range matches {
		if len(m) == 3 {
			o = append(o, m[2])
		}
	}
	return o
}

type Data struct {
	After    string `json:"after,omitempty"`
	Before   string `json:"before,omitempty"`
	Children []Link `json:"children,omitempty"`
	Dist     int    `json:"dist,omitempty"`
	Modhash  string `json:"modhash,omitempty"`
}

type LinkFlairRichText struct {
	E string `json:"e"`
	T string `json:"t"`
}

func (l Link) IsDailyDiscussion() bool {
	if len(l.Data.LinkFlairRichText) == 0 {
		return false
	}
	return l.Data.LinkFlairRichText[0].T == "Daily Discussion"
}
func (r Data) ChildIDs() (ids []string) {
	ids = make([]string, 0, len(r.Children))
	for _, val := range r.Children {
		ids = append(ids, val.Data.ID)
	}
	return ids
}

type RedditCommentResponse []Data

type ListingArray []RedditListing

type RedditListing struct {
	Kind string
	Data Data `json:"data"`
}

func (r RedditCommentResponse) AllChildren() (t []Link) {
	totalChildren := 0
	for _, v := range r {
		totalChildren += len(v.Children)
	}

	t = make([]Link, 0, totalChildren)

	for _, listing := range r {
		t = append(t, listing.Children...)
	}
	return
}

func (r ListingArray) AllChildren() (t []Link) {
	totalChildren := 0
	for _, v := range r {
		totalChildren += len(v.Data.Children)
	}

	t = make([]Link, 0, totalChildren)

	for _, listing := range r {
		t = append(t, listing.Data.Children...)
	}
	return
}

func (b Board) Subreddit() string {
	if len(b.Data.Children) == 0 {
		return ""
	}
	return b.Data.Children[0].Data.Subreddit
}

func (b Board) GetAllDiscussion() (t []T3Data, waitTime time.Duration, err error) {
	childIDs := b.Data.ChildIDs()

	for _, v := range childIDs {
		listings, er := GetCommentListing(b.Subreddit(), v, "?sort=top")

		if er != nil {
			err = er
		}
		for _, child := range listings.AllChildren() {
			t = append(t, child.Data.GetStockSymbols())
		}
		time.Sleep(time.Second + time.Second/4)

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

func GetBoard(boardName string) (t []T3Data, waitTime time.Duration, err error) {
	var board Board
	b, _, err := BrowserRequest(redditURL + fmt.Sprintf("/r/%s.json", boardName))

	if err != nil {
		return
	}

	err = json.Unmarshal(b, &board)
	if err != nil {
		return
	}
	for _, v := range board.Data.Children {
		t = append(t, v.Data.GetStockSymbols())
	}

	comments, waitTime, err := board.GetAllDiscussion()
	t = append(t, comments...)

	return
}

type T1 struct {
	Data T1Data
}

type T1Data struct {
	TotalAwardsReceived   int     `json:"total_awards_received,omitempty"`
	Ups                   uint    `json:"ups,omitempty"`
	LinkId                string  `json:"link_id,omitempty"`
	AuthorFullname        string  `json:"author_fullname,omitempty"`
	Id                    string  `json:"id,omitempty"`
	Gilded                uint    `json:"gilded,omitempty"`
	Author                string  `json:"author,omitempty"`
	ParentId              string  `json:"parent_id,omitempty"`
	Score                 int     `json:"score,omitempty"`
	SubredditId           string  `json:"subreddit_id,omitempty"`
	Body                  string  `json:"body,omitempty"`
	Downs                 uint    `json:"downs,omitempty"`
	BodyHtml              string  `json:"body_html,omitempty"`
	Stickied              bool    `json:"stickied,omitempty"`
	SubredditType         string  `json:"subreddit_type,omitempty"`
	Permalink             string  `json:"permalink,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Created               int64   `json:"created,omitempty"`
	Subreddit             string  `json:"subreddit,omitempty"`
	CreatedUtc            uint64  `json:"created_utc,omitempty"`
	SubredditNamePrefixed string  `json:"subreddit_name_prefixed,omitempty"`
	Controversiality      float64 `json:"controversiality,omitempty"`
	Depth                 uint16  `json:"depth,omitempty"`
}
