package repository

import (
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
	"github.com/susiraTH41/train-station/entities"

) 

type TrainStationRepository interface {
	GetStationNearMe(stationFilter *_trainStationModel.StationFilter) ([]*entities.TrainStation, error)
	GetStationNearMeOnPage(stationFilter *_trainStationModel.StationPaginateFilter) ([]*entities.TrainStation, error)
	CountStationNearMe(stationFilter *_trainStationModel.StationPaginateFilter) (int64, error)
	Creating(tempEntity *entities.Temp) (*entities.Temp, error)
	GetTemp() ([]*entities.Temp, error)
}