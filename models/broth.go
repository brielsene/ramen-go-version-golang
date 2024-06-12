package models

import "time"

type Broth struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt,omitempty"`
	ImageInactive string     `json:"imageInactive"`
	ImageActive   string     `json:"imageActive"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
}
