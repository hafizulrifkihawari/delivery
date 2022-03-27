package migration

import (
	"delivery/app/config"
	"delivery/constants"
	"delivery/domains/entities"
	"delivery/domains/models"
	"delivery/utils"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Migrate schema
func CreateSchema() {
	err := config.DBORM.AutoMigrate(
		&entities.Restaurant{},
		&entities.Menu{},
		&entities.RestaurantOpeningHour{},
		&entities.User{},
		&entities.UserPurchaseHistory{},
	)
	if err != nil {
		panic(err)
	}
}

func SeedData() {
	// seeder validation restaurant
	err := config.DBORM.First(&entities.Restaurant{}).Error
	if err == gorm.ErrRecordNotFound {
		seedRestaurant()
	}

	// seeder validation user
	err = config.DBORM.First(&entities.User{}).Error
	if err == gorm.ErrRecordNotFound {
		seedUser()
	}
}

func seedRestaurant() {
	var (
		restaurantSeeder []*models.RestaurantSeeder

		bulkInsertRestaurant []*entities.Restaurant
		bulkInsertMenu       []*entities.Menu
		bulkInsertOpening    []*entities.RestaurantOpeningHour
	)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(dir, "migration", "seeder", "restaurant_with_menu.json")

	// read json seeder
	seed, _ := ioutil.ReadFile(filePath)
	err = json.Unmarshal(seed, &restaurantSeeder)
	if err != nil {
		panic(err)
	}

	startID := 1
	for _, r := range restaurantSeeder {
		restaurant := &entities.Restaurant{
			ID:          uint(startID),
			Name:        r.RestaurantName,
			CashBalance: r.CashBalance,
		}

		for _, m := range r.Menu {
			bulkInsertMenu = append(bulkInsertMenu, &entities.Menu{
				RestaurantID: restaurant.ID,
				DishName:     m.DishName,
				Price:        m.Price,
			})
		}

		// parse opening hours
		splitted := strings.Split(r.OpeningHours, " / ")
		for _, t := range splitted {
			if t[len(t)-1:] == " " {
				t = t[:len(t)-1]
			}

			iv := strings.Split(t, " ")
			openTime := strings.Join(iv[len(iv)-5:len(iv)-3], " ")
			closeTime := strings.Join(iv[len(iv)-2:], " ")
			findLastIndex := strings.Index(t, openTime)
			dayTime := t[:findLastIndex-1]

			// remove non alphanumeric
			var replaceText string
			if strings.Contains(dayTime, " ") {
				replaceText = ""
			} else {
				replaceText = " "
			}

			removed := utils.ReplaceStringRegex(constants.AlphanumericSpace, dayTime, replaceText)
			listDays := strings.Fields(removed)
			for _, day := range listDays {
				bulkInsertOpening = append(bulkInsertOpening, &entities.RestaurantOpeningHour{
					RestaurantID: restaurant.ID,
					Day:          day,
					OpenAt:       utils.FormatStringTo24hourTime(openTime),
					CloseAt:      utils.FormatStringTo24hourTime(closeTime),
				})
			}
		}

		bulkInsertRestaurant = append(bulkInsertRestaurant, restaurant)
		startID += 1
	}

	insertionList := []interface{}{bulkInsertRestaurant, bulkInsertMenu, bulkInsertOpening}
	for i := range insertionList {
		bulkInsertion(insertionList[i])
	}
}

func seedUser() {
	var (
		userSeeder                   []*models.UserSeeder
		bulkInsertionUser            []*entities.User
		bulkInsertionPurchaseHistory []*entities.UserPurchaseHistory
	)

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := filepath.Join(dir, "migration", "seeder", "users_with_purchase_history.json")

	// read json seeder
	seed, _ := ioutil.ReadFile(filePath)
	err = json.Unmarshal(seed, &userSeeder)
	if err != nil {
		panic(err)
	}

	for _, u := range userSeeder {
		user := &entities.User{
			ID:          &u.ID,
			Name:        u.Name,
			CashBalance: u.CashBalance,
		}

		for _, h := range u.PurchaseHistory {
			td, err := time.Parse(constants.DateTime, h.TransactionDate)
			if err != nil {
				panic(err)
			}

			bulkInsertionPurchaseHistory = append(bulkInsertionPurchaseHistory, &entities.UserPurchaseHistory{
				UserID:            u.ID,
				DishName:          h.DishName,
				TransactionAmount: h.TransactionAmount,
				TransactionDate:   td,
			})
		}

		bulkInsertionUser = append(bulkInsertionUser, user)
	}

	insertionList := []interface{}{bulkInsertionUser, bulkInsertionPurchaseHistory}
	for i := range insertionList {
		bulkInsertion(insertionList[i])
	}
}

func bulkInsertion(data interface{}) {
	var defaultBatchSize = 500
	err := config.DBORM.CreateInBatches(data, defaultBatchSize).Error
	if err != nil {
		utils.PrintLog(err)
		panic(err)
	}
}

func InitMigration() {
	CreateSchema()
	SeedData()
}
