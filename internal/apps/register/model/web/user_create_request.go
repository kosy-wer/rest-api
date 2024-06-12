package web

// UserCreateRequest represents the request payload for creating a user
// swagger:model
type UserCreateRequest struct {
    // Name of the user
    // required: true
    // min length: 1
    // max length: 100
    Name string `json:"name"`
}
