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

type Restaurant struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	CashBalance float64 `json:"cash_balance"`
	Menu        []Menu
}

type RestaurantOpeningHour struct {
	ID           uint       `json:"id"`
	RestaurantID uint       `json:"restaurant_id"`
	Restaurant   Restaurant `json:"restaurant"`
	Day          string     `json:"day"`
	OpenAt       string     `json:"open_at"`
	CloseAt      string     `json:"close_at"`
}

type Menu struct {
	ID           uint       `json:"id"`
	RestaurantID uint       `json:"restaurant_id"`
	Restaurant   Restaurant `json:"restaurant"`
	DishName     string     `json:"dish_name"`
	Price        float64    `json:"price"`
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
	Day          string `json:"day"`
	Datetime     string `json:"datetime"`
	RestaurantID uint   `json:"restaurant_id"`
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
