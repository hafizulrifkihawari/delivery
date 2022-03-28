package logics

import (
	"delivery/app/config"
	"delivery/constants"
	"delivery/domains/models"
	"delivery/services/repositories"
	"delivery/utils"
	"fmt"
	"net/url"
	"strconv"
)

type IRestaurantService interface {
	FetchRestaurants(filterType string, param map[string]string) (interface{}, error)
}

type RestaurantService struct {
	restaurantRepository repositories.IRestaurantRepository
}

func InitRestaurantService(restaurantRepo repositories.IRestaurantRepository) *RestaurantService {
	if utils.IsNil(restaurantRepo) {
		restaurantRepo = repositories.InitRestaurantRepository(config.DBORM, config.DB)
	}

	service := RestaurantService{
		restaurantRepository: restaurantRepo,
	}
	return &service
}

func (service *RestaurantService) FetchRestaurants(filterType string, param map[string]string) (interface{}, error) {
	switch filterType {
	case "date":
		restaurants, err := service.ListRestaurantByFilterDatetime(param)
		return restaurants, err
	case "dish":
		restaurants, err := service.ListRestaurantByFilterDish(param)
		return restaurants, err
	case "search":
		restaurants, err := service.SearchRestaurant(param)
		return restaurants, err
	default:
		return nil, fmt.Errorf("%s filter type is not allowed", filterType)
	}
}

func (service *RestaurantService) ListRestaurantByFilterDatetime(param map[string]string) ([]*models.ListRestaurantResponse, error) {
	var (
		result     []*models.ListRestaurantResponse = []*models.ListRestaurantResponse{}
		filterArgs *models.FilterDate               = &models.FilterDate{}
	)

	epochTime, _ := strconv.Atoi(param["datetime"])

	parse := utils.ConvertEpochToTime(epochTime).UTC()
	weekday := parse.Weekday()

	filterArgs.Datetime = parse.Format(constants.MilitaryTime)
	filterArgs.Day = constants.DayMapping[int(weekday)]
	res, err := service.restaurantRepository.FilterByDatetime(filterArgs)
	if err != nil {
		go utils.PrintLog(err)
		return result, nil
	}

	for i := range res {
		restaurant := &models.ListRestaurantResponse{}
		_ = utils.AutoMap(res[i].Restaurant, &restaurant)

		restaurant.OpeningHour.Day = res[i].Day
		restaurant.OpeningHour.OpenAt = res[i].OpenAt
		restaurant.OpeningHour.CloseAt = res[i].CloseAt
		result = append(result, restaurant)
	}

	return result, nil
}

func (service *RestaurantService) ListRestaurantByFilterDish(param map[string]string) ([]*models.ListRestaurantResponse, error) {
	var (
		result     []*models.ListRestaurantResponse = []*models.ListRestaurantResponse{}
		filterArgs *models.FilterDish               = &models.FilterDish{}
	)

	// suffix reference:
	// gt -> greater than
	// lt -> less than
	// default operator set to lt, and/or if both value provided, gt will take priority
	if param["num_dishes_gt"] != "" {
		filterArgs.Operator = ">"
		filterArgs.NumDishes, _ = strconv.Atoi(param["num_dishes_gt"])
	} else {
		filterArgs.Operator = "<"
		filterArgs.NumDishes, _ = strconv.Atoi(param["num_dishes_lt"])
	}

	filterArgs.PriceStart, _ = strconv.ParseFloat(param["price_start"], 64)
	filterArgs.PriceEnd, _ = strconv.ParseFloat(param["price_end"], 64)
	filterArgs.Limit, _ = strconv.Atoi(param["limit"])

	res, err := service.restaurantRepository.FilterByNumDishesAndPrice(filterArgs)
	if err != nil {
		go utils.PrintLog(err)
		return result, nil
	}

	for i := range res {
		restaurant := &models.ListRestaurantResponse{}
		_ = utils.AutoMap(res[i], &restaurant)

		result = append(result, restaurant)
	}

	return result, nil
}

func (service *RestaurantService) SearchRestaurant(param map[string]string) ([]*models.ListRestaurantResponse, error) {
	var (
		result     []*models.ListRestaurantResponse = []*models.ListRestaurantResponse{}
		filterArgs *models.FilterText               = &models.FilterText{}
	)

	if param["term"] != "" {
		filterArgs.Term = url.QueryEscape(param["term"])
	}
	res, err := service.restaurantRepository.FilterByTextSearch(filterArgs)
	if err != nil {
		go utils.PrintLog(err)
		return result, nil
	}

	_ = utils.AutoMap(res, &result)

	return result, nil
}
