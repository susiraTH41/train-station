package repository

import (
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
	"github.com/susiraTH41/train-station/entities"

) 

type TrainStationRepository interface {
	GetStationNearMe(stationFilter *_trainStationModel.StationFilter) ([]*entities.TrainStation, error)
	GetStationNearMeOnPage(stationFilter *_trainStationModel.StationPaginateFilter) ([]*entities.TrainStation, error)
}