package trello

import (
	"testing"
)

func TestGetList(t *testing.T) {
	list := testList(t)
	if list.Name != "To Do Soon" {
		t.Errorf("Title incorrect. Got '%s'", list.Name)
	}
}

func TestGetListsOnBoard(t *testing.T) {
	board := testBoard(t)
	board.client.BaseURL = mockResponse("lists", "board-lists-api-example.json").URL
	lists, err := board.GetLists(Defaults)
	if err != nil {
		t.Fatal(err)
	}

	if len(lists) != 3 {
		t.Errorf("Expected 1 list, got %d", len(lists))
	}
}

// Utility function to get the standard case Client.GetList() response
//
func testList(t *testing.T) *List {
	c := NewClient("user", "pass")
	c.BaseURL = mockResponse("lists", "list-api-example.json").URL
	list, err := c.GetList("4eea4ff", Defaults)
	if err != nil {
		t.Fatal(err)
	}
	return list
}