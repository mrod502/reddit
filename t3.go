package reddit

import (
	"encoding/json"
	"fmt"
	"time"
)

//T3Data - a top-level post (as opposed to T1 replies)
type T3Data struct {
	AllAwardings        []Awarding          `json:"all_awardings,omitempty"`
	Author              string              `json:"author,omitempty"`
	AuthorFullname      string              `json:"author_fullname,omitempty"`
	Body                string              `json:"body,omitempty"`
	Controversiality    float64             `json:"controversiality,omitempty"`
	Created             float64             `json:"created,omitempty"`
	Depth               uint16              `json:"depth,omitempty"`
	Downs               uint                `json:"downs,omitempty"`
	Gilded              uint                `json:"gilded,omitempty"`
	Id                  string              `json:"id,omitempty"`
	LinkFlairRichText   []LinkFlairRichText `json:"link_flair_richtext,omitempty"`
	LinkId              string              `json:"link_id,omitempty"`
	Name                string              `json:"name,omitempty"`
	ParentId            string              `json:"parent_id,omitempty"`
	Permalink           string              `json:"permalink,omitempty"`
	Score               int                 `json:"score,omitempty"`
	Selftext            string              `json:"selftext,omitempty"`
	Stickied            bool                `json:"stickied,omitempty"`
	Subreddit           string              `json:"subreddit,omitempty"`
	SubredditId         string              `json:"subreddit_id,omitempty"`
	SubredditType       string              `json:"subreddit_type,omitempty"`
	Symbols             []string            `json:"symbols,omitempty"`
	Title               string              `json:"title,omitempty"`
	TotalAwardsReceived int                 `json:"total_awards_received,omitempty"`
	Ups                 uint                `json:"ups,omitempty"`
	UpvoteRatio         float64             `json:"upvote_ratio,omitempty"`
}

func (t T3Data) Compact() CompactT3 {

	return CompactT3{
		Title:   t.Title,
		Text:    t.Selftext,
		Symbols: GetStockSymbols(t.Title + "\n" + t.Selftext),
		Created: time.Unix(int64(t.Created), 0),
		Ups:     t.Ups,
	}
}

func (t T3Data) GetStockSymbols() T3Data {
	t.Symbols = GetStockSymbols(t.Selftext + "\n" + t.Title + "\n" + t.Body)
	return t
}

func (t T3Data) GetReplies() (d []T3Data, err error) {
	b, _, err := BrowserRequest(fmt.Sprintf("%sr/%s/comments/%s.json", redditBase, t.Subreddit, t.Id))
	var res RedditCommentResponse
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(b, &res)
	res.AllChildren()
	return
}
