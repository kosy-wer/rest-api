package web

// ErrorResponse represents an error response
// swagger:response errorResponse
type ErrorResponse struct {
    // The HTTP status code
    // example: 400
    Code int `json:"code"`

    // The status message corresponding to the status code
    // example: "Bad Request"
    Status string `json:"status"`

    // The error details
    // example: {"message": "Invalid request payload"}
    Error struct {
        Message string `json:"message"`
    } `json:"error"`
}
