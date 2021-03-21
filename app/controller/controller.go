package controller

import (
	"encoding/json"
	"fmt"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/app/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	GetTraceIp(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
	GetStatistics(w http.ResponseWriter, r *http.Request)
}

type GenericController struct {
	TraceIpService   service.TraceIpServiceInterface
	StatisticService service.StatisticServiceInterface
}

func NewController(traceIpService service.TraceIpServiceInterface, statisticService service.StatisticServiceInterface) Controller {
	return &GenericController{
		TraceIpService:   traceIpService,
		StatisticService: statisticService,
	}
}

func (controller *GenericController) GetTraceIp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ip := params["ip"]
	fmt.Println("Params: ", ip)

	result, err := controller.TraceIpService.GetTraceIp(ip)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
	} else {
		json.NewEncoder(w).Encode(result)
		statisticItem := model.StatisticItem{
			CountryName: result.Country,
			Distance:    result.Distance,
		}
		controller.StatisticService.IncrementScore(statisticItem)
	}
}

func (controller *GenericController) GetStatistics(w http.ResponseWriter, r *http.Request) {
	result, err := controller.StatisticService.GetScores()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		if result == (model.Statistic{}) {
			fmt.Fprintf(w, "empty statistics")
		} else {
			json.NewEncoder(w).Encode(result)
		}
	}
}

// Home
func (controller *GenericController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Api")
}
