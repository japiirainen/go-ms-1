package model

type Cat struct {
	// postgres id for a cat
	ID int `json:"id"`
	// breed of the cat
	Breed string `json:"breed"`
	// color of the cat
	Color string `json:"color"`
	// gender of the cat
	Gender Gender `json:"gender"`
	// owner of the cat (relation to owner by owbers id)
	Owner string `json:"owner"`
}

type NewCat struct {
	ID     int    `json:"id"`
	Breed  string `json:"breed"`
	Color  string `json:"color"`
	Gender Gender `json:"gender"`
	Owner  string `json:"owner"`
}
