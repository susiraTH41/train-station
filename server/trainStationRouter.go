package server

import (
	"github.com/labstack/echo/v4"
	_trainStationController "github.com/susiraTH41/train-station/pkg/trainStation/controller"
	_trainStationRepository "github.com/susiraTH41/train-station/pkg/trainStation/repository"
	_trainStationService "github.com/susiraTH41/train-station/pkg/trainStation/service"
)

func (s *echoServer) initTrainStationRouter(middleware echo.MiddlewareFunc){
	router := s.app.Group("/v1/train-station")

	trainStationRepository := _trainStationRepository.NewTrainStationRepositoryImpl(s.db, s.app.Logger)
	trainStationService := _trainStationService.NewTrainStationServiceImpl(trainStationRepository)
	trainStationController := _trainStationController.NewTrainStationControllerImpl(trainStationService)

	router.GET("/getstationclosetome", trainStationController.GetStationNearMe ,middleware)
	router.GET("/getstationclosetomeonpage", trainStationController.GetStationNearMeOnPage ,middleware)

}
