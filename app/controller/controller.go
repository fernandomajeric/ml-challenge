package controller

import (
	"encoding/json"
	"fmt"
	"github.com/fernandomajeric/ml-challenge/app/model"
	"github.com/fernandomajeric/ml-challenge/app/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(result)
	statisticItem := model.StatisticItem{
		CountryName: result.Country,
		Distance:    result.Distance,
	}
	err = controller.StatisticService.IncrementScore(statisticItem)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (controller *GenericController) GetStatistics(w http.ResponseWriter, r *http.Request) {
	result, err := controller.StatisticService.GetScores()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
