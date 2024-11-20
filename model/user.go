package model

type Users struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Username string `json:"username,omitempty"`
	Address  string `json:"address,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Phone    string `json:"phone,omitempty" validate:"required,min=10"`
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty" validate:"required"`
}
