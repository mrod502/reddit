package reddit

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	var client = NewClient(time.Minute)

	board, err := client.GetSubreddit("wallstreetbets")
	if err != nil {
		t.Fatal(err)
	}

	replies, _ := board.GetPostReplies(board.Data.Children[0].Data)
	fmt.Printf("%d\n", len(replies))
}
