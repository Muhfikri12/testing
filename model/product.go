package model

type Products struct {
	ID                 uint        `json:"id" validate:"required"`
	Name               string      `json:"name"`
	ImageUrl           string      `json:"image_url"`
	Price              int         `json:"price"`
	PriceAfterDiscount int         `json:"price_after_discount,omitempty"`
	Discount           int         `json:"discount"`
	Description        string      `json:"description"`
	Status             string      `json:"status,omitempty"`
	Variants           *[]Variant  `json:"variants,omitempty"`
	Categories         *Categories `json:"categories,omitempty"`
	Images             *Images     `json:"images,omitempty"`
	Previews           *Previews   `json:"previews,omitempty"`
	Checkouts          *Checkouts  `json:"checkouts,omitempty"`
	Timestamps         *Basic      `json:"timestamps,omitempty"`
}

type Categories struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Images struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Variant struct {
	Size     string `json:"size,omitempty"`
	Color    string `json:"color,omitempty"`
	Stocks   int    `json:"stock,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}
