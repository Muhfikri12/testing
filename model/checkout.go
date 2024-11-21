package model

type Checkouts struct {
	ID          uint        `json:"id,omitempty"`
	Users       *Users      `json:"users,omitempty"`
	Products    *[]Products `json:"products,omitempty"`
	UserID      string      `json:"user_id,omitempty"`
	TotalAmount float64     `json:"total_amount,omitempty"`
	Payment     string      `json:"payment,omitempty"`
	TotalSold   int         `json:"total_sold,omitempty"`
	Timestamps  *Basic      `json:"timestamp,omitempty"`
}
