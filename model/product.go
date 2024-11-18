package model

type Products struct {
	ID          uint        `json:"id" validate:"required"`
	Name        string      `json:"name"`
	ImageUrl    string      `json:"image_url"`
	Price       string      `json:"price"`
	Discount    string      `json:"discount"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Categories  *Categories `json:"categories,omitempty"`
	Images      *Images     `json:"images,omitempty"`
	Previews    *Previews   `json:"previews,omitempty"`
	Checkouts   *Checkouts  `json:"checkouts,omitempty"`
	Timestamps  *Basic      `json:"timestamps,omitempty"`
}

type Categories struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Images struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
