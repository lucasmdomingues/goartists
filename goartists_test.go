package goartists

import (
	"testing"
)

const appID = "123"

func TestSearch(t *testing.T) {
	service := NewService(appID)

	_, err := service.Search("Megadeth")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetEvents(t *testing.T) {
	service := NewService(appID)

	artist, err := service.Search("Megadeth")
	if err != nil {
		t.Error(err)
		return
	}

	if err = service.GetEvents(artist); err != nil {
		t.Error(err)
		return
	}
}
