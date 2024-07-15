package service

import (
	"time"

	"github.com/susiraTH41/train-station/entities"
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
	_trainStationRepository "github.com/susiraTH41/train-station/pkg/trainStation/repository"
)

type trainStationServiceImpl struct {
	trainStationRepository _trainStationRepository.TrainStationRepository
}

func NewTrainStationServiceImpl(
	trainStationRepository _trainStationRepository.TrainStationRepository,
) TrainStationService {
	return &trainStationServiceImpl{trainStationRepository}
}


func (s *trainStationServiceImpl) GetStationNearMe(stationFilter *_trainStationModel.StationFilter) ([]*_trainStationModel.Station, error){
	stationList, err := s.trainStationRepository.GetStationNearMe(stationFilter)
	if err != nil {
		return nil, err
	}

	stationModelList := make([]*_trainStationModel.Station, 0)

	for _, station := range stationList {
		stationModelList = append(stationModelList, station.ToStationModel())
	}


	
	return stationModelList, nil
}

func (s *trainStationServiceImpl) GetStationNearMeOnPage(stationFilter *_trainStationModel.StationPaginateFilter) (*_trainStationModel.StationOnPage, error){

	stationList, err := s.trainStationRepository.GetStationNearMeOnPage(stationFilter)
	if err != nil {
		return nil, err
	}

	count, err := s.trainStationRepository.CountStationNearMe(stationFilter)
	if err != nil {
		return nil, err
	}


	stationModelList := make([]*_trainStationModel.Station, 0)

	for _, station := range stationList {
		stationModelList = append(stationModelList, station.ToStationModel())
	}

	totalPage := count / stationFilter.Size
	if totalPage % stationFilter.Size != 0 {
		totalPage++
	}
	
	return &_trainStationModel.StationOnPage{
				Station: stationModelList,
				Page:    stationFilter.Page,
				Size:    totalPage,
			}, err
}

func (s *trainStationServiceImpl) Creating(itema string) (*entities.Temp, error) {

	item := &entities.Temp{
		CreatedAt: time.Now(),
		TempoNow: itema,
	}

	tempEntity, err := s.trainStationRepository.Creating(item)
	if err != nil {
		return  nil,err
	}

	return tempEntity, err
}

func (s *trainStationServiceImpl) GetTemp() ([]*entities.Temp, error){
	stationList, err := s.trainStationRepository.GetTemp()
	if err != nil {
		return nil, err
	}
	
	return stationList, nil
}