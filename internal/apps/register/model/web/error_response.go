package web

// ErrorResponse represents an error response
// swagger:response
type ErrorResponse struct {
    // HTTP status code
    Code int `json:"code"`
    // Status message
    Status string `json:"status"`
    // Error details
    Error string `json:"error"`
}
