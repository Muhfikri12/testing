package model

type Previews struct {
	ID             uint     `json:"id,omitempty" validate:"required"`
	Description    string   `json:"description,omitempty"`
	Rating         *float64 `json:"rating,omitempty"`
	TotalReviewers int      `json:"total_reviewers,omitempty"`
}
