package web

// CategoryUpdateRequest represents the request payload for updating a category
// swagger:model
type CategoryUpdateRequest struct {
    // ID of the category
    // required: true
    Id int `validate:"required" json:"id"`
    // Updated name of the category
    // required: true
    // min length: 1
    // max length: 200
    Name string `validate:"required,max=200,min=1" json:"name"`
}
