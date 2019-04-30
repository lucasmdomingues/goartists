package goartists

type Artist struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	URL                string `json:"url"`
	ImageURL           string `json:"image_url"`
	ThumbURL           string `json:"thumb_url"`
	FacebookPageURL    string `json:"facebook_page_url"`
	MbID               string `json:"mbid"`
	TrackerCount       int64  `json:"tracker_count"`
	UpComingEventCount int64  `json:"upcoming_event_count"`
}

type Event struct {
	ID             string   `json:"id"`
	ArtistID       string   `json:"artist_id"`
	URL            string   `json:"url"`
	OnSaleDateTime string   `json:"on_sale_datetime"`
	Datetime       string   `json:"datetime"`
	Description    string   `json:"description"`
	Venue          *Venue   `json:"venue"`
	Offers         []*Offer `json:"offers"`
	Lineup         []string `json:"lineup"`
}

type Venue struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Country   string `json:"country"`
}

type Offer struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

type BandsintownError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
