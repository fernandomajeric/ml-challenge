package service

import (
	"fmt"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Mock of Repository
type repositoryMock struct {
	mock.Mock
}

func (mock *repositoryMock) Increment(statist model.StatisticItem) error {
	fmt.Println("Mocked Statistic - Increment notification function")
	fmt.Printf("Value passed in: %s\n", statist.CountryName)
	_ = mock.Called(statist)
	return nil
}

func (mock *repositoryMock) GetScores() map[string]model.StatisticCore {
	fmt.Println("Mocked Statistic - Increment notification function")
	args := mock.Called()
	return args.Get(0).(map[string]model.StatisticCore)
}

func TestGetScoresShouldCallRepository(t *testing.T) {
	// Arrange
	const countryCode string = "CA"
	scores := make(map[string]model.StatisticCore)
	scores[countryCode] = model.StatisticCore{
		Country:  countryCode,
		Distance: 100000,
		HitCount: 0,
	}
	statisticMock := new(repositoryMock)
	statisticMock.On("GetScores").Return(scores)
	sut := NewStatisticService(statisticMock)
	// Action
	_, _ = sut.GetScores()
	// Assert
	statisticMock.AssertExpectations(t)
}

func TestGetScoresWithScoreShouldGetScores(t *testing.T) {
	// Arrange
	const countryCode string = "CA"
	scores := make(map[string]model.StatisticCore)
	scores[countryCode] = model.StatisticCore{
		Country:  countryCode,
		Distance: 100000,
		HitCount: 0,
	}
	statisticExpected := model.Statistic{
		MinScore:        scores[countryCode],
		MaxScore:        scores[countryCode],
		AverageDistance: 0,
	}
	statisticMock := new(repositoryMock)
	statisticMock.On("GetScores").Return(scores)
	sut := NewStatisticService(statisticMock)
	// Action
	statistic, _ := sut.GetScores()
	// Assert
	statisticMock.AssertExpectations(t)
	assert.EqualValues(t, statisticExpected, statistic)
}

func TestIncrementScoreShouldCallRepository(t *testing.T) {
	// Arrange
	scoreItem:= model.StatisticItem{
		CountryName: "Canada",
		Distance:    10000,
	}
	statisticMock := new(repositoryMock)
	statisticMock.On("Increment", scoreItem)
	sut := NewStatisticService(statisticMock)
	// Action
	_ = sut.IncrementScore(scoreItem)
	// Assert
	statisticMock.AssertCalled(t,"Increment", scoreItem)
}
