package common

// HTTPResponse common data response for handling REST data and status
type HTTPResponse struct {
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Length     uint64      `json:"length,omitempty"`
	HTTPstatus int         `json:"-"`
}
