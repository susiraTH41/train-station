package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/susiraTH41/train-station/databases"
)

type trainStationRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewTrainStationRepositoryImpl(db databases.Database, logger echo.Logger) TrainStationRepository {
	return &trainStationRepositoryImpl{db, logger}
}
