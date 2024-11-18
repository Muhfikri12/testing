package model

type Checkouts struct {
	ID          uint    `json:"id,omitempty"`
	UserID      string  `json:"user_id,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	Payment     string  `json:"payment,omitempty"`
	TotalBuyers int     `json:"total_buyers,omitempty"`
	Timestamps  *Basic  `json:"timestamp,omitempty"`
}
