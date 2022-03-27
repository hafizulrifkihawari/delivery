package logics

import (
	"delivery/app/config"
	"delivery/constants"
	"delivery/domains/models"
	"delivery/services/repositories"
	"delivery/utils"
	"strconv"
)

type IRestaurantService interface {
	ListRestaurantByFilter(param map[string]string) ([]*models.ListRestaurantResponse, error)
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

func (service *RestaurantService) ListRestaurantByFilter(param map[string]string) ([]*models.ListRestaurantResponse, error) {
	var (
		result     []*models.ListRestaurantResponse = []*models.ListRestaurantResponse{}
		filterArgs *models.FilterArgs               = &models.FilterArgs{}
		// err        error
	)

	if param["datetime"] != "" {
		epochTime, _ := strconv.Atoi(param["datetime"])

		parse := utils.ConvertEpochToTime(epochTime).UTC()
		weekday := parse.Weekday()

		filterArgs.Datetime = parse.Format(constants.MilitaryTime)
		filterArgs.Day = constants.DayMapping[int(weekday)]
		res, err := service.restaurantRepository.FetchRestaurantByDatetime(filterArgs)
		if err != nil {
			go utils.PrintLog(err)
			return result, err
		}

		for i := range res {
			restaurant := &models.ListRestaurantResponse{}
			_ = utils.AutoMap(res[i].Restaurant, &restaurant)

			restaurant.OpeningHour.Day = res[i].Day
			restaurant.OpeningHour.OpenAt = res[i].OpenAt
			restaurant.OpeningHour.CloseAt = res[i].CloseAt
			result = append(result, restaurant)
		}
	}

	return result, nil
}
