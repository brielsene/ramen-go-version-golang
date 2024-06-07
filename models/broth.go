package models

import "gorm.io/gorm"

type Broth struct {
	gorm.Model
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`

	// private Long id;
	// private String imageInactive;
	// private String imageActive;
	// private String name;

	// private String description;
	// private BigDecimal price;
}
