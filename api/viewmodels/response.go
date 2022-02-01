package viewmodels

// Response -
type Response struct {
	Success   bool        `json:"success"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
	Error     interface{} `json:"error"`
}
