package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/susiraTH41/train-station/config"
	"github.com/susiraTH41/train-station/databases"
	"github.com/susiraTH41/train-station/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	TrainStationData(db)

	tx := db.Connect().Begin()

	AddStationFromAPI(tx, conf.LinkAPI.GetStationUrl)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

}


func TrainStationData(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.TrainStation{})

}

func AddStationFromAPI(tx *gorm.DB, url string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(body)

	if err != nil {
		panic(err.Error())
	}
	var stationJsonList []entities.TrainStation

	json.Unmarshal(body, &stationJsonList)


	tx.CreateInBatches(stationJsonList, len(stationJsonList))
}

