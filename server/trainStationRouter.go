package server

import (
	_trainStationController "github.com/susiraTH41/train-station/pkg/trainStation/controller"
	_trainStationRepository "github.com/susiraTH41/train-station/pkg/trainStation/repository"
	_trainStationService "github.com/susiraTH41/train-station/pkg/trainStation/service"
)

func (s *echoServer) initTrainStationRouter() {
	router := s.app.Group("/v1/train-station")

	trainStationRepository := _trainStationRepository.NewTrainStationRepositoryImpl(s.db, s.app.Logger)
	trainStationService := _trainStationService.NewTrainStationServiceImpl(trainStationRepository)
	trainStationController := _trainStationController.NewTrainStationControllerImpl(trainStationService)

	// router.GET("/getstationclosetome", trainStationController.GetStationNearMe )
	// router.GET("/getstationclosetomeonpage", trainStationController.GetStationNearMeOnPage )
	router.POST("/temp", trainStationController.Creating)
	router.GET("/gettemp", trainStationController.GetTemp)

}
