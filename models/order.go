package models

type Order struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`

	// 	"description": "Salt and Chasu Ramen",
	//   "image": "https://tech.redventures.com.br/icons/ramen/ramenChasu.png"
}
