package models

type UserSeeder struct {
	ID              uint                  `gorm:"primary_key" json:"id"`
	Name            string                `json:"name"`
	CashBalance     float64               `json:"cashBalance"`
	PurchaseHistory []UserPurchaseHistory `json:"purchaseHistory"`
}

type UserPurchaseHistory struct {
	ID                uint    `gorm:"primary_key" json:"id"`
	RestaurantName    string  `json:"restaurantName"`
	DishName          string  `json:"dishName"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionDate   string  `json:"transactionDate"`
}
