package service

import(
	_trainStationRepository "github.com/susiraTH41/train-station/pkg/trainStation/repository"
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"

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

func (s *trainStationServiceImpl) GetStationNearMeOnPage(stationFilter *_trainStationModel.StationPaginateFilter) ([]*_trainStationModel.Station, error){

	stationList, err := s.trainStationRepository.GetStationNearMeOnPage(stationFilter)
	if err != nil {
		return nil, err
	}

	stationModelList := make([]*_trainStationModel.Station, 0)

	for _, station := range stationList {
		stationModelList = append(stationModelList, station.ToStationModel())
	}


	
	return stationModelList, nil
}
