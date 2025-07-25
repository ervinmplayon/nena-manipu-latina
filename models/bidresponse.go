package models

type BidResponse struct {
	Creative string  `json:"creative"`
	CPM      float64 `json:"cpm"`
	Bidder   string  `json:"bidder"`
}
