package model

import "time"

type Promotions struct {
	ID         uint      `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Subtitle   string    `json:"subtitle,omitempty"`
	ImageUrl   string    `json:"image_url,omitempty"`
	PathUrl    string    `json:"path_url,omitempty"`
	StartDate  time.Time `json:"start_date,omitempty"`
	EndDate    string    `json:"end_date,omitempty"`
	Timestamps *Basic    `json:"timestamps,omitempty"`
}
