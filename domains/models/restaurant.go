package models

type RestaurantSeeder struct {
	CashBalance    float64      `json:"cashBalance"`
	Menu           []MenuSeeder `json:"menu"`
	OpeningHours   string       `json:"openingHours"`
	RestaurantName string       `json:"restaurantName"`
}

type MenuSeeder struct {
	DishName string  `json:"dishName"`
	Price    float64 `json:"price"`
}
