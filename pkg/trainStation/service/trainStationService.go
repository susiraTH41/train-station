package service

import (
	"github.com/susiraTH41/train-station/entities"
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
)

 type TrainStationService interface{
	GetStationNearMe(stationFilter *_trainStationModel.StationFilter) ([]*_trainStationModel.Station, error)
	GetStationNearMeOnPage(stationFilter *_trainStationModel.StationPaginateFilter) (*_trainStationModel.StationOnPage, error)
	Creating(item string)  (*entities.Temp, error)
	GetTemp() ([]*entities.Temp, error)
}