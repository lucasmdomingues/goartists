package goartists

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const bandSintownPrefix = "https://rest.bandsintown.com"

func SearchArtist(appID, artistName string) (*Artist, error) {

	url := fmt.Sprintf("%s/artists/%s?app_id=%s", bandSintownPrefix, artistName, appID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 || strings.Contains(string(body), "error") {
		return nil, newBandsintownError(body)
	}

	if len(string(body)) == 2 {
		return nil, fmt.Errorf("%s", "No artist returned")
	}

	var artist *Artist

	err = json.Unmarshal(body, &artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (artist *Artist) GetEvents(appID string) ([]*Event, error) {

	url := fmt.Sprintf("%s/artists/%s/events?app_id=%s", bandSintownPrefix, artist.Name, appID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 || strings.Contains(string(body), "error") {
		return nil, newBandsintownError(body)
	}

	if strings.TrimSpace(string(body)) == "[]" {
		return nil, fmt.Errorf("%s", "No events returned")
	}

	var events []*Event

	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func newBandsintownError(error []byte) error {

	var bandsintownError *BandsintownError

	err := json.Unmarshal(error, &bandsintownError)
	if err != nil {
		return err
	}

	if bandsintownError.Message != "" {
		err = fmt.Errorf("%s", bandsintownError.Message)
	} else if bandsintownError.Error != "" {
		err = fmt.Errorf("%s", bandsintownError.Error)
	}

	return fmt.Errorf("%s", err)
}
