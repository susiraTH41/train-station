package entities


type TrainStationData struct {
	ID              uint64  `gorm:"primaryKey;type:bigint;"`
	StationCode     uint64  `gorm:"not null;"`
	Name            string  `gorm:"type:varchar(128);unique;not null;"`
	EnName          string  `gorm:"type:varchar(128);not null;"`
	ThShort         string  `gorm:"type:varchar(64);not null;"`
	EnShort         string  `gorm:"type:varchar(64);not null;"`
	Chname          string  `gorm:"type:varchar(64);not null;"`
	Controldivision int     `gorm:"not null;"`
	ExactKm         int     `gorm:"not null;"`
	ExactDistance   int     `gorm:"not null;"`
	Km              int     `gorm:"not null;"`
	Class           int     `gorm:"not null;"`
	Lat             float64 `gorm:"not null;"`
	Long            float64 `gorm:"not null;"`
	Active          int     `gorm:"not null;"`
	Giveway         int     `gorm:"not null;"`
	DualTrack       int     `gorm:"not null;"`
	Comment         string  `gorm:"type:varchar(64);not null;"`
}