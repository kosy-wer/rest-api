package web

// WebResponse is a generic response structure used for API responses.
// swagger:response webResponse
type WebResponse struct {
    // The HTTP status code
    // example: 200
    Code int `json:"code"`

    // The status message corresponding to the status code
    // example: "OK"
    Status string `json:"status"`

    // The data payload of the response
    // example: {"email": "john@gmail.com", "name": "John Doe"}
    Data interface{} `json:"data"`
}
