package entities

import "time"

type User struct {
	ID          *uint   `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	CashBalance float64 `json:"cash_balance"`
}

type UserPurchaseHistory struct {
	ID                uint      `gorm:"primary_key" json:"id"`
	DishName          string    `json:"dish_name"`
	UserID            uint      `json:"user_id"`
	User              User      `gorm:"<-:false,foreignKey:UserID" json:"user"`
	TransactionAmount float64   `json:"transaction_amount"`
	TransactionDate   time.Time `json:"transaction_date"`
}
