package entities

type TrainStation struct {
	ID              int64     `json:"id" gorm:"primaryKey;type:bigint;"`
	StationCode     int64     `json:"station_code"  gorm:"not null;"`
	Name            string  `json:"name"  gorm:"type:varchar(128);unique;not null;"`
	EnName          string  `json:"en_name"  gorm:"type:varchar(128);not null;"`
	ThShort         string  `json:"th_short"gorm:"type:varchar(64);not null;"`
	EnShort         string  `json:"en_short"gorm:"type:varchar(64);not null;"`
	Chname          string  `json:"chname"gorm:"type:varchar(64);"`
	Controldivision int64     `json:"controldivision" type:"integer" gorm:"not null;"`
	ExactKm         int64     `json:"exact_km" type:"integer" gorm:"not null; "`
	ExactDistance   int64     `json:"exact_distance" type:"integer" gorm:"not null; ;"`
	Km              int64     `json:"km" type:"integer" gorm:"not null; "`
	Class           int64     `json:"class"gorm:"not null;"`
	Lat             float64 `json:"lat"gorm:"not null;type:decimal(10,6);"`
	Long            float64  `json:"long"gorm:"not null;type:decimal(10,6);"`
	Active          bool     `json:"active"gorm:""`
	Giveway         bool     `json:"giveway"gorm:"not null;"`
	DualTrack       bool     `json:"dual_track"gorm:"not null;"`
	Comment         string  `json:"comment"gorm:"type:varchar(64);"`
}