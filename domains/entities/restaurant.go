package entities

type Restaurant struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	CashBalance float64 `json:"cash_balance"`
}

type RestaurantOpeningHour struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	RestaurantID uint       `json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"<-:false,foreignKey:RestaurantID" json:"restaurant"`
	Day          string     `json:"day"`
	OpenAt       string     `json:"open_at"`
	CloseAt      string     `json:"close_at"`
}

type Menu struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	RestaurantID uint       `json:"restaurant_id"`
	Restaurant   Restaurant `gorm:"<-:false,foreignKey:RestaurantID" json:"restaurant"`
	DishName     string     `json:"dish_name"`
	Price        float64    `json:"price"`
}
