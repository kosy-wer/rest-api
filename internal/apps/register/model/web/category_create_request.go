package web

// CategoryCreateRequest represents the request payload for creating a category
// swagger:model
type CategoryCreateRequest struct {
    // Name of the category
    // required: true
    // min length: 1
    // max length: 100
    Name string `json:"name"`
}
