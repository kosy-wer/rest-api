package web

// UserUpdateRequest represents the request payload for updating a user
// swagger:model
type UserUpdateRequest struct {
    // ID of the user
    // required: true
    Id int `validate:"required" json:"id"`
    // Updated name of the user
    // required: true
    // min length: 1
    // max length: 200
    Name string `validate:"required,max=200,min=1" json:"name"`
}
