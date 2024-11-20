package model

type Cart struct {
	ID            uint      `json:"id,omitempty"`
	UserID        uint      `json:"user_id,omitempty"`
	Products      *Products `json:"products,omitempty"`
	TotalProducts uint      `json:"total_products,omitempty"`
}
