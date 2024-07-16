package entities

import "time"

type Temp struct {
	TempoNow    string  `json:"en_short"gorm:"type:varchar(64);"`
	Humidity    string  `json:"en_short"gorm:"type:varchar(64);"`
	CreatedAt   time.Time `json:"createdat" gorm:"not null;autoCreateTime;"`
}
