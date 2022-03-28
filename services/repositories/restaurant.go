package repositories

import (
	"database/sql"
	"delivery/domains/entities"
	"delivery/domains/models"
	"delivery/utils"
	"fmt"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	connORM *gorm.DB
	connDB  *sql.DB
}

type IRestaurantRepository interface {
	FilterByDatetime(filterArgs *models.FilterDate) ([]*entities.RestaurantOpeningHour, error)
	FilterByNumDishesAndPrice(filterArgs *models.FilterDish) ([]*entities.Restaurant, error)
	FilterByTextSearch(filterArgs *models.FilterText) ([]*entities.RestaurantSearch, error)
}

func InitRestaurantRepository(connORM *gorm.DB, connDB *sql.DB) *RestaurantRepository {
	return &RestaurantRepository{
		connORM: connORM,
		connDB:  connDB,
	}
}

func (repo *RestaurantRepository) FilterByDatetime(filterArgs *models.FilterDate) ([]*entities.RestaurantOpeningHour, error) {
	var result []*entities.RestaurantOpeningHour
	err := repo.connORM.Joins("Restaurant").Where("day = ? AND ? BETWEEN open_at AND close_at", filterArgs.Day, filterArgs.Datetime).Find(&result).Error
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}
	return result, nil
}

func (repo *RestaurantRepository) FilterByNumDishesAndPrice(filterArgs *models.FilterDish) ([]*entities.Restaurant, error) {
	var (
		rows   *sql.Rows
		err    error
		result = []*entities.Restaurant{}
	)

	query := fmt.Sprintf(`SELECT r.id, r.name
	FROM restaurant r
	WHERE id IN (
		SELECT DISTINCT(restaurant_id) 
		FROM "menu" 
		WHERE price BETWEEN $1 AND $2
	) AND (
		SELECT COUNT(-1) FROM menu m WHERE r.id = m.restaurant_id GROUP BY m.restaurant_id
	) %s $3
	ORDER BY r.name
	LIMIT $4;`, filterArgs.Operator)
	rows, err = repo.connDB.Query(query, filterArgs.PriceStart, filterArgs.PriceEnd, filterArgs.NumDishes, filterArgs.Limit)
	if err != nil {
		go utils.PrintLog(err)
		if err == sql.ErrNoRows {
			return result, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		restaurant := &entities.Restaurant{}
		err = rows.Scan(&restaurant.ID, &restaurant.Name)
		if err != nil {
			go utils.PrintLog(err)
			return nil, err
		}
		result = append(result, restaurant)
	}
	return result, nil
}

func (repo *RestaurantRepository) FilterByTextSearch(filterArgs *models.FilterText) ([]*entities.RestaurantSearch, error) {
	var result []*entities.RestaurantSearch
	err := repo.connORM.Where("search_text @@ plainto_tsquery(?)", filterArgs.Term).Find(&result).Error
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}
	return result, nil
}
