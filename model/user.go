package model

type Users struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}
