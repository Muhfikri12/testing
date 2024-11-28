package model

type Users struct {
	ID              int        `json:"id,omitempty"`
	Name            string     `json:"name,omitempty" validate:"required"`
	Username        string     `json:"username,omitempty"`
	Address         *Addresses `json:"addresses,omitempty"`
	Email           string     `json:"email,omitempty" validate:"required,email"`
	Phone           string     `json:"phone,omitempty" validate:"required,min=10,alphanum"`
	Token           string     `json:"token,omitempty"`
	TotalShipping   int        `json:"total_shipping,omitempty"`
	CurrentPassword *string    `json:"current_password,omitempty"`
	NewPassword     string     `json:"new_password,omitempty"`
	ConfirmPassword string     `json:"confirm_password,omitempty" validate:"eqfield=NewPassword"`
}

type Login struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty" binding:"required,email"`
	Phone    string `json:"phone,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	Token    string `json:"token,omitempty"`
}

type Addresses struct {
	ID       uint   `json:"id,omitempty"`
	Address  string `json:"address,omitempty" validate:"required"`
	Longlat  string `json:"longlat,omitempty"`
	Distanse string `json:"distanse,omitempty"`
	Status   string `json:"status,omitempty"`
}

type Register struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Phone    string `json:"phone,omitempty" validate:"required,min=10,alphanum"`
	Password string `json:"password,omitempty" validate:"required"`
}
