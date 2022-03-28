package models

type User struct {
	ID          *uint   `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	CashBalance float64 `json:"cash_balance"`
}

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

type UserPurchaseRequest struct {
	MenuID int `json:"menu_id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
}

type UserPurchaseResponse struct {
	Name              string  `json:"name"`
	RestaurantName    string  `json:"restaurant_name"`
	DishName          string  `json:"dish_name"`
	TransactionAmount float64 `json:"transaction_amount"`
	CashBalance       float64 `json:"cash_balance"`
}
