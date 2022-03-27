package repositories

import (
	"database/sql"
	"delivery/domains/entities"
	"delivery/domains/models"
	"delivery/utils"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	connORM *gorm.DB
	connDB  *sql.DB
}

type IRestaurantRepository interface {
	FetchRestaurantByDatetime(filterArgs *models.FilterArgs) ([]*entities.RestaurantOpeningHour, error)
}

func InitRestaurantRepository(connORM *gorm.DB, connDB *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{
		connORM: connORM,
		connDB:  connDB,
	}
}

func (repo *RestaurantRepository) FetchRestaurantByDatetime(filterArgs *models.FilterArgs) ([]*entities.RestaurantOpeningHour, error) {
	var result []*entities.RestaurantOpeningHour
	err := repo.connORM.Joins("Restaurant").Where("day = ? AND ? BETWEEN open_at AND close_at", filterArgs.Day, filterArgs.Datetime).Find(&result).Error
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}
	return result, nil
}
