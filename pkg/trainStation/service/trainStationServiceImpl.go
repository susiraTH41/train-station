package service

import(
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