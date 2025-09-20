package web

// UserCreateRequest represents the request payload for creating a new user
// swagger:model
type UserCreateRequest struct {
    // First name of the user
    // required: true
    FirstName string `validate:"required,min=1,max=100" json:"first_name"`

    // Last name of the user
    // required: false
    LastName string `validate:"max=100" json:"last_name"`

    // Email of the user
    // required: true
    Email string `validate:"required,email" json:"email"`

    // Password of the user
    // required: true
    Password string `validate:"required,min=6,max=100" json:"password"`
}

