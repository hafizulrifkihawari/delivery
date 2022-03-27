package models

type ListRestaurantResponse struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	CashBalance float64     `json:"cash_balance"`
	OpeningHour OpeningHour `json:"opening_hour"`
}

type OpeningHour struct {
	Day     string `json:"day"`
	OpenAt  string `json:"open_at"`
	CloseAt string `json:"close_at"`
}

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

type FilterArgs struct {
	Day      string `json:"day"`
	Datetime string `json:"datetime"`
}
