package web

// UserCreateRequest represents the request payload for updating a user
// swagger:model
type UserCreateRequest struct {
    // Email of the user
    // required: true
    Email string `validate:"required,email" json:"email"`
    // Updated name of the user
    // required: true
    // min length: 1
    // max length: 200
    Name string `validate:"required,max=200,min=1" json:"name"`
}
