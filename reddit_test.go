package reddit

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	var client = NewClient()

	err := client.AddSub("wallstreetbets")
	if err != nil {
		t.Fatal(err)
	}
	board, err := GetSub("wallstreetbets")
	if err != nil {
		t.Fatal(err)
	}

	replies, _ := board.GetPostReplies(board.Data.Children[0].Data)
	fmt.Printf("%d\n", len(replies))
}
