package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/susiraTH41/train-station/databases"

	"github.com/susiraTH41/train-station/entities"
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
)

type trainStationRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewTrainStationRepositoryImpl(db databases.Database, logger echo.Logger) TrainStationRepository {
	return &trainStationRepositoryImpl{db, logger}
	
}

var distance_calculation = `SQRT(POWER(train_stations.lat - ?,2) +  POWER(train_stations.lng  - ?,2))`


func (r *trainStationRepositoryImpl) GetStationNearMe(stationFilter *_trainStationModel.StationFilter) ([]*entities.TrainStation, error){
	trainStation := make([]*entities.TrainStation, 0)
	numLimit := stationFilter.Count

	err := r.db.Connect().Table("train_stations").
        Select("train_stations.id, train_stations.en_name, train_stations.name, train_stations.lat, train_stations.lng, "+
            distance_calculation+" as distance\n", stationFilter.Lat, stationFilter.Long).
		Where("active = ?", 1).
        Order("distance ASC").
        Limit(numLimit).Scan(&trainStation).Error		

	if err != nil {
		r.logger.Errorf("Failed to train stations: %s", err.Error())
		return nil, err

	}

	return trainStation, nil
}


func (r *trainStationRepositoryImpl) GetStationNearMeOnPage(stationFilter *_trainStationModel.StationPaginateFilter) ([]*entities.TrainStation, error){
	trainStation := make([]*entities.TrainStation, 0)
	
	offset := int((stationFilter.Page - 1) * stationFilter.Size)
	limit := int(stationFilter.Size)

	err := r.db.Connect().Table("train_stations").
        Select("train_stations.id, train_stations.en_name, train_stations.name, train_stations.lat, train_stations.lng, "+
            distance_calculation+" as distance\n", stationFilter.Lat, stationFilter.Long).
		Where("active = ?", 1).
        Order("distance ASC").
		Offset(offset).Limit(int(limit)).
        Scan(&trainStation).Error	

	

	if err != nil {
		r.logger.Errorf("Failed to train stations on page: %s", err.Error())
		return nil, err
	}

	return trainStation, nil

}


func (r *trainStationRepositoryImpl) CountStationNearMe(stationFilter *_trainStationModel.StationPaginateFilter) (int64, error){
	var count int64

	err := r.db.Connect().Table("train_stations").
        Select("train_stations.id, train_stations.en_name, train_stations.name, train_stations.lat, train_stations.lng, "+
            distance_calculation+" as distance\n", stationFilter.Lat, stationFilter.Long).
		Where("active = ?", 1).
        Order("distance ASC").
		Count(&count).Error	

	if err != nil {
		r.logger.Errorf("Failed to train stations on page: %s", err.Error())
		return -1, err
	}

	return count, nil

}

func (r *trainStationRepositoryImpl) Creating(tempEntity *entities.Temp) (*entities.Temp, error){

	insertedTemp := new(entities.Temp)

	if err := r.db.Connect().Create(tempEntity).Scan(insertedTemp).Error; err != nil {
		r.logger.Error("Creating player failed", err.Error())
		return nil, err
	}

	return insertedTemp, nil
	

}//GetTemp() (*entities.Temp, error)


func (r *trainStationRepositoryImpl) GetTemp() ([]*entities.Temp, error){

	query := r.db.Connect().Model(&entities.Temp{})

	items := make([]*entities.Temp, 0)


	if err := query.Order("created_at DESC").Limit(1).Find(&items).Error; err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, err
	}

	return items, nil
}


