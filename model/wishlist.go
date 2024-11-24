package model

type Wishlists struct {
	ID         int   `json:"id"`
	ProductID  int   `json:"product_id"`
	UserID     int   `json:"user_id"`
	Timestamps Basic `json:"timestamp,omitempty"`
}
