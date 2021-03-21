package service

import (
	"errors"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/app/repository"
	"math"
)

type StatisticServiceInterface interface {
	IncrementScore(score model.StatisticItem) error
	GetScores() (statistic model.Statistic, err error)
}

type StatisticService struct{
	Repository repository.StatisticRepositoryInterface
}

func NewStatisticService(repository repository.StatisticRepositoryInterface) *StatisticService {
	return &StatisticService{Repository: repository}
}

func (service *StatisticService) IncrementScore(score model.StatisticItem) error {
	return service.Repository.Increment(score)
}

func (service *StatisticService) GetScores() (statistic model.Statistic, err error) {
	scores := service.Repository.GetScores()
	values := getValues(scores)
	minScore, _ := getMinValue(values)
	maxScore, _ := getMaxValue(values)

	return model.Statistic{
		MinScore:        minScore,
		MaxScore:        maxScore,
		AverageDistance: getAverage(values),
	}, nil
}

func getValues(keyValues map[string]model.StatisticCore) (values []model.StatisticCore) {
	var list []model.StatisticCore
	for _, s := range keyValues {
		list = append(list, s)
	}
	return list
}

func getMinValue(values []model.StatisticCore) (min model.StatisticCore, e error) {
	if len(values) == 0 {
		return model.StatisticCore{}, errors.New("array empty")
	}

	min = values[0]
	for _, v := range values {
		if v.HitCount < min.HitCount {
			min = v
		}
	}
	return min, nil
}

func getMaxValue(values []model.StatisticCore) (max model.StatisticCore, e error) {
	if len(values) == 0 {
		return model.StatisticCore{}, errors.New("array empty")
	}

	max = values[len(values)-1]
	for _, v := range values {
		if v.HitCount > max.HitCount {
			max = v
		}
	}
	return max, nil
}

func getAverage(scores []model.StatisticCore) float64 {
	var totalHitsCount = 0.0
	var totalDistance = 0.0

	for _, s := range scores {
		totalDistance = totalDistance + float64(s.HitCount)*s.Distance
		totalHitsCount = totalHitsCount + float64(s.HitCount)
	}

	if totalHitsCount == 0.0 {
		return 0.0
	}

	return math.Round(totalDistance / totalHitsCount)
}
