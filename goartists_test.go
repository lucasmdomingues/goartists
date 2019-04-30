package goartists

import (
	"testing"
)

const appID = ""

func TestSearchArtist(t *testing.T) {

	artist, err := SearchArtist(appID, "Megadeth")
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", artist)
}

func TestGetEvents(t *testing.T) {

	artist, err := SearchArtist(appID, "Megadeth")
	if err != nil {
		t.Error(err)
		return
	}

	events, err := artist.GetEvents(appID)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", events[0].Venue)
}
