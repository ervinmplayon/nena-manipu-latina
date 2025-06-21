package models

type AdResponse struct {
	AdMarkup string  `json:"ad_markup"`
	CPM      float64 `json:"cpm"`
	Bidder   string  `json:"bidder"`
}
