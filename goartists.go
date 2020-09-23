package goartists

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://rest.bandsintown.com"

type service struct {
	AppID string
}

type Service interface {
	Search(name string) (*Artist, error)
	GetEvents(artist *Artist) error
}

func NewService(appID string) Service {
	return &service{
		AppID: appID,
	}
}

func (s *service) Search(name string) (*Artist, error) {
	url := fmt.Sprintf("%s/artists/%s?app_id=%s", URL, name, s.AppID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newBandsintownError(body)
	}

	if string(body) == `""` {
		return nil, errors.New("Oops, artists not found")
	}

	var artist *Artist

	err = json.Unmarshal(body, &artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (s *service) GetEvents(artist *Artist) error {
	if artist == nil {
		return errors.New("Oops, artist invalid")
	}

	url := fmt.Sprintf("%s/artists/%s/events?app_id=%s", URL, artist.Name, s.AppID)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return newBandsintownError(body)
	}

	return json.Unmarshal(body, &artist.Events)
}

func newBandsintownError(errByte []byte) error {
	var bandsintownError *BandsintownError

	err := json.Unmarshal(errByte, &bandsintownError)
	if err != nil {
		return err
	}

	err = fmt.Errorf("Oops, %s", bandsintownError.Message)

	return err
}
