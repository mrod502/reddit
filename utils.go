package reddit

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

const (
	redditBase = "https://www.reddit.com/"
)

var (
	mdLinkRegex = regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`)
	symbolRegex = regexp.MustCompile(`(\$|#)([A-Z0-9]+)`)
)

var (
	ErrNotFound      = errors.New("item not found")
	ErrTypeAssertion = errors.New("unable to assert type")
)

type CompactT3 struct {
	Title   string    `json:"title,omitempty"`
	Text    string    `json:"text,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Symbols []string  `json:"symbols,omitempty"`
	Links   []string  `json:"links,omitempty"`
	Ups     uint
}

// BrowserRequest -- pretend to be a browser so we can get comments
func BrowserRequest(url string) (b []byte, rh http.Header, err error) {

	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Set("upgrade-insecure-requests", "1")
	r.Header.Set("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`)
	r.Header.Set("accept-language", "en-US,en;q=0.9,hr-HR;q=0.8,hr;q=0.7,ru-RU;q=0.6,ru;q=0.5")
	r.Header.Set("scheme", "https")
	r.Header.Set("authority", "www.reddit.com")

	cli := http.DefaultClient

	resp, err := cli.Do(r)
	if err != nil {
		return
	}
	rh = resp.Header

	b, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return
}

// EnableCORS - enable cross-origin requests
func EnableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "privatekey")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,POST,HEAD,DELETE,PUT")
}

func BoardUri(b string) string {
	return redditBase + "r" + b
}
