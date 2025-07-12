package models

type BidRequest struct {
	RequestID   string `json:"request_id"`
	PublisherID string `json:"publisher_id"`
	AdUnit      string `json:"ad_unit"`
	DeviceIP    string `json:"device_ip"`
	UserAgent   string `json:"user_agent"`
	// TODO: Geo, UA, App info, Device, etc.
}
