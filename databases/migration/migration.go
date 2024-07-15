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

	// TrainStationData(db)

	tx := db.Connect().Begin()
	tempMigration(tx)


	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	// AddStationFromAPI(tx, conf.LinkAPI.GetStationUrl)

	// if err := tx.Commit().Error; err != nil {
	// 	tx.Rollback()
	// 	panic(err)
	// }

	// ch := time.NewTicker(1 * time.Second)

	// go func() {
	// 	for range ch.C {
	// 		log.Printf("Tick at: %v\n", time.Now())

	// 	}
	// }()

	// time.Sleep(5 * time.Second)
    // ch.Stop()

	// c := cron.New()

	// c.AddFunc("@every 1s",      
	// func() { AddStationFromAPI(tx, conf.LinkAPI.GetStationUrl) })
	// c.Start()

	// select {}

}


func TrainStationData(db databases.Database) {
	db.Connect().Migrator().CreateTable(&entities.TrainStation{})

}

func tempMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Temp{})
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


	tx.Commit().UpdateColumns(stationJsonList)
	fmt.Println("done") 
}

