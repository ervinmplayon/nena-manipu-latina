package models

type AdRequest struct {
	RequestID   string `json:"request_id"`
	PublisherID string `json:"publisher_id"`
	AdUnit      string `json:"ad_unit"`
	DeviceIP    string `json:"device_ip"`
	// TODO: Geo, UA, App info, Device, etc.
}
