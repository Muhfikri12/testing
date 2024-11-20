package model

type Wishlists struct {
	ID         int   `json:"id"`
	ProductID  int   `json:"product_id" validate:"required,gte=1"`
	UserID     int   `json:"user_id" validate:"required,gte=1"`
	Timestamps Basic `json:"timestamp,omitempty"`
}
