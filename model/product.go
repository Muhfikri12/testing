package model

type Products struct {
	Users              *Users      `json:"users,omitempty"`
	ID                 uint        `json:"id,omitempty" validate:"required"`
	Name               string      `json:"name,omitempty"`
	ImageUrl           string      `json:"image_url,omitempty"`
	Price              int         `json:"price,omitempty"`
	PriceAfterDiscount int         `json:"price_after_discount,omitempty"`
	Amount             int         `json:"amount,omitempty"`
	Discount           int         `json:"discount,omitempty"`
	Qty                int         `json:"qty,omitempty"`
	Description        string      `json:"description,omitempty"`
	Status             string      `json:"status,omitempty"`
	Size               string      `json:"size,omitempty"`
	Color              string      `json:"color,omitempty"`
	Variants           *[]Variant  `json:"variants,omitempty"`
	Categories         *Categories `json:"categories,omitempty"`
	Previews           *Previews   `json:"previews,omitempty"`
	Checkouts          *Checkouts  `json:"checkouts,omitempty"`
	Timestamps         *Basic      `json:"timestamps,omitempty"`
}

type Categories struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Variant struct {
	Size     string  `json:"size,omitempty"`
	Color    string  `json:"color,omitempty"`
	Stocks   int     `json:"stock,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
}
