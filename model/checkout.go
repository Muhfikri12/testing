package model

type Checkouts struct {
	ID            uint        `json:"id,omitempty"`
	Users         *Users      `json:"users,omitempty"`
	UserID        string      `json:"user_id,omitempty"`
	Payment       string      `json:"payment,omitempty"`
	PaymentStatus string      `json:"payment_status,omitempty"`
	TotalSold     int         `json:"total_sold,omitempty"`
	Timestamps    *Basic      `json:"timestamp,omitempty"`
	Products      *[]Products `json:"products,omitempty"`
	ShippingCost  int         `json:"shipping_cost,omitempty"`
	TotalAmount   int         `json:"total_amount,omitempty"`
}
