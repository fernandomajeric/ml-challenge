package app

import (
	"github.com/fernandomajeric/ml-challenge/app/controller"
	"github.com/fernandomajeric/ml-challenge/app/http/rest"
	"github.com/fernandomajeric/ml-challenge/app/repository"
	"github.com/fernandomajeric/ml-challenge/app/service"
	"github.com/fernandomajeric/ml-challenge/config"
	"github.com/gorilla/mux"
	muxlogrus "github.com/pytimer/mux-logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type App interface {
	// Start http server
	Start(serverPort string)
}

type ApiApplication struct {
	config *config.GeneralConfig
}

// New : build new ApiApplication
func New(configFilePaths ...string) App {
	return &ApiApplication{
		config: config.LoadConfig(configFilePaths...),
	}
}

// Start serve http server
func (app *ApiApplication) Start(serverPort string) {
	// init new handler
	myRouter := mux.NewRouter().StrictSlash(true)

	//add logger middleware
	myRouter.Use(muxlogrus.NewLogger().Middleware)

	//Get Services
	traceIpService, statisticService := getServices()

	//Get Controller
	controller := controller.NewController(traceIpService, statisticService)

	//Api Routing map
	myRouter.HandleFunc("/", controller.Home)
	myRouter.HandleFunc("/trace-ip/{ip}", controller.GetTraceIp).Methods(http.MethodGet)
	myRouter.HandleFunc("/statistics", controller.GetStatistics).Methods(http.MethodGet)

	log.Info("Server Started at http://localhost:" + serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, myRouter))
}

func getServices() (m service.TraceIpServiceInterface, s service.StatisticServiceInterface) {
	traceIpService := service.NewTraceIpService(rest.NewGeoLocalizationRestClient(rest.Client),
		rest.NewCountryRestClient(rest.Client),
		rest.NewCurrencyRateRestClient(rest.Client))

	statisticService := service.NewStatisticService(repository.NewStatisticRepository())
	return traceIpService, statisticService
}
