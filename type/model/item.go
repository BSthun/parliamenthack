package model

type ScrapedItem struct {
	ImageUrl *string `json:"imageUrl"`
	Name     *string `json:"name"`
	Price    *string `json:"price"`
}
