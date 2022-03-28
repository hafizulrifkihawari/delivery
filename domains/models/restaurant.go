package models

type ListRestaurantResponse struct {
	ID             uint         `json:"id"`
	Name           string       `json:"name,omitempty"`
	RestaurantName string       `json:"restaurant_name,omitempty"`
	DishName       string       `json:"dish_name,omitempty"`
	OpeningHour    *OpeningHour `json:"opening_hour,omitempty"`
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

type FilterDate struct {
	Day      string `json:"day"`
	Datetime string `json:"datetime"`
}

type FilterText struct {
	Term string `json:"term"`
}

type FilterDish struct {
	Operator    string
	NumDishes   int
	NumDishesGT int     `json:"num_dishes_gt"`
	NumDishesLT int     `json:"num_dishes_lt"`
	PriceStart  float64 `json:"price_start"`
	PriceEnd    float64 `json:"price_end"`
	Limit       int     `json:"limit"`
}
