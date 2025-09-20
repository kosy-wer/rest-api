package web

type UserUpdateRequest struct {
    // Email of the user
    Email     string `validate:"required,email" json:"email"`

    // Updated first name of the user
    FirstName string `validate:"required,max=100,min=1" json:"first_name"`

    // Updated last name of the user
    LastName  string `validate:"required,max=100,min=1" json:"last_name"`

    // Optional: if password can be updated
    Password  string `validate:"omitempty,min=6" json:"password"`
}

