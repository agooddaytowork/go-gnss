package main

import (
	"fmt"
	"github.com/geoscienceaustralia/go-rtcm/rtcm3"
	"github.com/umeat/go-gnss/cmd/database/models"
	"github.com/umeat/go-gnss/cmd/database/util"
	"github.com/umeat/go-ntrip/ntrip"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("sqlite3", "database/test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Observation{})
	db.AutoMigrate(&models.SatelliteData{})
	db.AutoMigrate(&models.SignalData{})

	client, err := ntrip.NewClient("https://streams.auscors.geops.team/TEST00AUS0")
	resp, err := client.Connect()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode, err)
	}

	scanner := rtcm3.NewScanner(resp.Body)
	for frame, err := scanner.NextFrame(); err == nil; frame, err = scanner.NextFrame() {
		switch frame.MessageNumber() {
		case 1077, 1087, 1097, 1107, 1117, 1127:
			obs, _ := util.ObservationMsm7(rtcm3.DeserializeMessageMsm7(frame.Payload))
			db.Create(&obs)
		}
	}
	panic(err)
}
