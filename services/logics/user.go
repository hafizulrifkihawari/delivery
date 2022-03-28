package logics

import (
	"delivery/app/config"
	"delivery/constants"
	"delivery/domains/entities"
	"delivery/domains/models"
	"delivery/services/repositories"
	"delivery/utils"
	"errors"

	"gorm.io/gorm"
)

type IUserService interface {
	PurchaseDish(purchaseReq *models.UserPurchaseRequest) (*models.UserPurchaseResponse, error)
}

type UserService struct {
	restaurantLogic IRestaurantService
	userRepository  repositories.IUserRepository
	dbTrx           *gorm.DB
}

func InitUserService(userRepo repositories.IUserRepository, restaurantLogic IRestaurantService, db *gorm.DB, useTransaction bool) *UserService {
	if utils.IsNil(db) {
		db = config.DBORM
		if useTransaction {
			db = db.Begin()
		}
	}
	if utils.IsNil(userRepo) {
		userRepo = repositories.InitUserRepository(db, nil)
	}
	if utils.IsNil(restaurantLogic) {
		restaurantLogic = InitRestaurantService(nil, db)
	}

	service := UserService{
		restaurantLogic: restaurantLogic,
		userRepository:  userRepo,
		dbTrx:           db,
	}
	return &service
}

func (service *UserService) PurchaseDish(purchaseReq *models.UserPurchaseRequest) (*models.UserPurchaseResponse, error) {
	var (
		timeNow = constants.TimeNow
		weekday = timeNow.Weekday()
		user    *models.User
	)

	// fetch user data
	userData, err := service.userRepository.GetUserByID(uint(purchaseReq.UserID))
	if err != nil {
		go utils.PrintLog(err)
		return nil, errors.New(constants.ErrorUserNotFound)
	}
	_ = utils.AutoMap(userData, &user)

	// Fetch dish data
	dish, err := service.restaurantLogic.GetDishByID(uint(purchaseReq.MenuID))
	if err != nil {
		go utils.PrintLog(err)
		return nil, errors.New(constants.ErrorMenuNotFound)
	}

	// validation restaurant still open
	filterDate := &models.FilterDate{
		Day:          constants.DayMapping[int(weekday)],
		Datetime:     timeNow.Format(constants.MilitaryTime),
		RestaurantID: dish.Restaurant.ID,
	}
	_, err = service.restaurantLogic.CheckRestaurantByOpeningHour(filterDate)
	if err != nil {
		go utils.PrintLog(err)
		return nil, errors.New(constants.ErrorRestaurantClosed)
	}

	// validation user cash balance
	if user.CashBalance-dish.Price < 0 {
		return nil, errors.New(constants.ErrorInsufficientFund)
	}

	resp, err := service.ProcessPurchasing(user, dish)
	if err != nil {
		service.dbTrx.Rollback()
		go utils.PrintLog(err)
		return nil, err
	}

	service.dbTrx.Commit()

	return resp, nil
}

func (service *UserService) ProcessPurchasing(user *models.User, dish *models.Menu) (*models.UserPurchaseResponse, error) {
	var (
		// calculation balance
		userBalance       = user.CashBalance - dish.Price
		restaurantBalance = dish.Restaurant.CashBalance + dish.Price
	)

	// insert into user history
	userHistory := &entities.UserPurchaseHistory{
		UserID:            *user.ID,
		DishName:          dish.DishName,
		TransactionAmount: dish.Price,
		TransactionDate:   constants.TimeNow,
	}

	err := service.userRepository.CreateUserPurchaseHistory(userHistory)
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}

	// update user balance
	err = service.userRepository.UpdateUserBalance(*user.ID, userBalance)
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}

	// update restaurant balance
	err = service.restaurantLogic.UpdateRestaurantBalance(dish.Restaurant.ID, restaurantBalance)
	if err != nil {
		go utils.PrintLog(err)
		return nil, err
	}

	response := &models.UserPurchaseResponse{
		Name:              user.Name,
		RestaurantName:    dish.Restaurant.Name,
		DishName:          dish.DishName,
		TransactionAmount: dish.Price,
		CashBalance:       userBalance,
	}

	return response, nil
}
