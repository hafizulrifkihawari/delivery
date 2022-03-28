package repositories

import (
	"database/sql"
	"delivery/domains/entities"
	"delivery/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	connORM *gorm.DB
	connDB  *sql.DB
}

type IUserRepository interface {
	GetUserByID(id uint) (*entities.User, error)
	CreateUserPurchaseHistory(userHistory *entities.UserPurchaseHistory) error
	UpdateUserBalance(userID uint, cashBalance float64) error
}

func InitUserRepository(connORM *gorm.DB, connDB *sql.DB) *UserRepository {
	return &UserRepository{
		connORM: connORM,
		connDB:  connDB,
	}
}

func (repo *UserRepository) GetUserByID(id uint) (*entities.User, error) {
	var result *entities.User
	err := repo.connORM.First(&result, id).Error
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}
	return result, nil
}

func (repo *UserRepository) CreateUserPurchaseHistory(userHistory *entities.UserPurchaseHistory) error {
	err := repo.connORM.Create(&userHistory).Error
	if err != nil {
		go utils.PrintLog(err)
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateUserBalance(userID uint, cashBalance float64) error {
	err := repo.connORM.Model(&entities.User{}).Where("id = ?", userID).Update("cash_balance", cashBalance).Error
	if err != nil {
		go utils.PrintLog(err)
		return err
	}
	return nil
}
