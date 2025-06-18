package models

type AdResponse struct {
	AdMarkup string `json:"ad_markup"`
	CPM      string `json:"cpm"`
	Bidder   string `json:"bidder"`
}
