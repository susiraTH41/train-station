package controller

import (

	_trainStationService "github.com/susiraTH41/train-station/pkg/trainStation/service"
)

type trainStationControllerImpl struct {
	trainStationService _trainStationService.TrainStationService
}

func NewTrainStationControllerImpl( 
	trainStationService _trainStationService.TrainStationService ,
) TrainStationController {
	return &trainStationControllerImpl{trainStationService}
}
