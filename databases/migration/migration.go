package main

import (
	"encoding/json"
	"fmt"
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

	AddStationFromAPI(tx, conf.MigrationData.GetStationUrl)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}

}

func TrainStationData(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.TrainStation{})

}

func AddStationFromAPI(tx *gorm.DB, url string) {
	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var stationJsonList []entities.TrainStation

	json.Unmarshal(body, &stationJsonList)

	// var newStationList []entities.TrainStation

	// for _, p := range stationJsonList {
	// 	var newStation entities.TrainStation	
	// 	newStation.ID = p.ID
	// 	newStation.StationCode = p.StationCode
	// 	newStation.Name = p.Name
	// 	newStation.EnName = p.EnName
	// 	newStation.ThShort = p.ThShort
	// 	newStation.EnShort = p.EnShort
	// 	newStation.Chname = p.Chname
	// 	newStation.Controldivision = p.Controldivision
	// 	newStation.ExactKm = p.ExactKm
	// 	newStation.ExactDistance = p.ExactDistance
	// 	newStation.Km = p.Km
	// 	newStation.Class = p.Class
	// 	newStation.Lat = p.Lat
	// 	newStation.Long = p.Long
	// 	newStation.Active = p.Active
	// 	newStation.Giveway = p.Giveway
	// 	newStation.DualTrack = p.DualTrack
	// 	newStation.Comment = p.Comment

		// fmt.Println(stationJsonList)
	// 	newStationList = append(newStationList, newStation)
	// }
	

	// var station []entities.TrainStation
	// err = json.NewDecoder(res.Body).Decode(&station)
	// if err != nil {
	//     fmt.Println("Error decoding JSON:", err)
	//     return
	// }

	// Now you have the parsed JSON data in the "data" variable
	// fmt.Println("Fetched data:", station[0])


	tx.CreateInBatches(stationJsonList, len(stationJsonList))
}
