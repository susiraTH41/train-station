package entities

import (
	_trainStationModel "github.com/susiraTH41/train-station/pkg/trainStation/model"
)


type TrainStation struct {
	ID              int64     `json:"id" gorm:"primaryKey;type:bigint;not null;"`
	StationCode     int64     `json:"station_code"`
	Name            string    `json:"name"  gorm:"type:varchar(128);"`
	EnName          string  `json:"en_name"  gorm:"type:varchar(128);"`
	ThShort         string  `json:"th_short"gorm:"type:varchar(64);"`
	EnShort         string  `json:"en_short"gorm:"type:varchar(64);"`
	Chname          string  `json:"chname"gorm:"type:varchar(64);"`
	Controldivision int     `json:"controldivision" type:"integer"`
	ExactKm         int     `json:"exact_km" type:"integer"`
	ExactDistance   int     `json:"exact_distance" type:"integer"`
	Km              int     `json:"km" type:"integer" `
	Class           int     `json:"class"`
	Lat 			float64 `json:"lat" gorm:"type:decimal(10,6);`
	Lng 			float64 `json:"long" gorm:"type:decimal(10,6);`
	Active          int8     `json:"active"`
	Giveway         int8     `json:"giveway"`
	DualTrack       int8     `json:"dual_track"`
	Comment         string  `json:"comment"gorm:"type:varchar(64);"`
	Distance    	float64  	`json:"distance"`

}




func (i *TrainStation) ToStationModel() *_trainStationModel.Station {

	return &_trainStationModel.Station{
		Name:   i.Name,
		Lat: 	i.Lat,
		Long:   i.Lng,
		EnName: i.EnName,
		Distance: i.Distance,

	}
}	
