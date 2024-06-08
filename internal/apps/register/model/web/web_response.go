package web

// WebResponse represents the response format for web requests
// swagger:response
type WebResponse struct {
    // HTTP status code
    Code int `json:"code"`
    // Status message
    Status string `json:"status"`
    // Response data
    Data interface{} `json:"data"`
}
